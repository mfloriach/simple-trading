package trade

import (
	"context"
)

type cacheService struct {
	service TradeRepository
}

func NewTradeCacheService() TradeService {
	return &cacheService{service: NewTradeCacheRepository()}
}

func (s cacheService) Put(ctx context.Context, userID int, t TradeAction) error {
	return s.service.Set(ctx, Trade{
		Symbol: t.Symbol,
		Amount: t.Amount,
		UserID: userID,
		Action: ActionPut,
	})
}

func (s cacheService) Call(ctx context.Context, userID int, t TradeAction) error {
	return s.service.Set(ctx, Trade{
		Symbol: t.Symbol,
		Amount: t.Amount,
		UserID: userID,
		Action: ActionCall,
	})
}

func (s cacheService) TakeProfit(ctx context.Context, userID int, t TradeAction) error {
	return s.service.Set(ctx, Trade{
		Symbol: t.Symbol,
		Amount: t.Amount,
		UserID: userID,
		Action: ActionTakeProfit,
	})
}

func (s cacheService) StopLose(ctx context.Context, userID int, t TradeAction) error {
	return s.service.Set(ctx, Trade{
		Symbol: t.Symbol,
		Amount: t.Amount,
		UserID: userID,
		Action: ActionStopLose,
	})
}
