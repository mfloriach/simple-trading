package graphs

import (
	"context"
	"fmt"
	"sync"
	"time"

	"math/rand"

	"example.com/grpc-todo/server/internal/trade"
)

type ApiGatewayService interface {
	Subscribe() (int, chan Candle)
	UnSubscribe(pos int)
	PullCandle(ctx context.Context)
}

type Candle struct {
	Symbol string
	High   int
	Open   int
	Low    int
	Close  int
	Time   time.Time
}

type apiGatewayService struct {
	sync.RWMutex
	symbol           string
	candleRepository GraphRepository
	tradeRepository  trade.TradeRepositoryAdmin
	subscribers      []chan<- Candle
}

func NewApiGatewayService(symbol string, tradeRepository trade.TradeRepositoryAdmin, candleRepository GraphRepository) *apiGatewayService {
	return &apiGatewayService{
		symbol:           symbol,
		subscribers:      make([]chan<- Candle, 0),
		tradeRepository:  tradeRepository,
		candleRepository: candleRepository,
	}
}

func (s *apiGatewayService) Subscribe() (int, chan Candle) {
	ch := make(chan Candle, 1)

	s.Lock()
	defer s.Unlock()
	s.subscribers = append(s.subscribers, ch)

	return len(s.subscribers) - 1, ch
}

func (s *apiGatewayService) UnSubscribe(pos int) {
	s.Lock()
	defer s.Unlock()

	s.subscribers = append(s.subscribers[:pos], s.subscribers[pos+1:]...)
}

func (s *apiGatewayService) PullCandle(ctx context.Context) {
	var (
		ticker    = time.NewTicker(time.Second)
		lastPrice int
	)

	for {
		price := rand.Intn(150)

		go s.checkStopLoseAndTakeProfit(ctx, price, lastPrice)

		lastPrice = price

		if err := s.candleRepository.Insert(ctx, s.symbol, price); err != nil {
			fmt.Println("Error inserting data into database", err)
		}

		candle := Candle{
			Symbol: s.symbol,
			Open:   price,
			Close:  price - 30,
			High:   price + 30,
			Low:    price - 50,
			Time:   time.Now(),
		}

		s.Lock()
		for _, subscriber := range s.subscribers {
			subscriber <- candle
		}
		s.Unlock()

		<-ticker.C
	}
}

func (s *apiGatewayService) checkStopLoseAndTakeProfit(ctx context.Context, price, lastPrice int) {
	if price > lastPrice {
		trades, err := s.tradeRepository.Get(ctx, s.symbol, trade.ActionTakeProfit)
		if err != nil {
			fmt.Println("Error getting take profit trade: ", err)
		}

		for _, t := range trades {
			if price > int(t.Amount) {
				fmt.Println("Take profit", t)
				if err := s.tradeRepository.Delete(ctx, t.ID); err != nil {
					fmt.Println("Error deleting take profit trade", err)
				}
			}
		}

		return
	}

	trades, err := s.tradeRepository.Get(ctx, s.symbol, trade.ActionStopLose)
	if err != nil {
		fmt.Println("Error getting stop lose trade: ", err)
	}

	for _, t := range trades {
		if price < int(t.Amount) {
			fmt.Println("stop loss", t)
			if err := s.tradeRepository.Delete(ctx, t.ID); err != nil {
				fmt.Println("Error deleting stop lose trade", err)
			}
		}
	}

}
