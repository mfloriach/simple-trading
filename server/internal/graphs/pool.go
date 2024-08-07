package graphs

import (
	"context"

	"example.com/grpc-todo/server/internal/trade"
)

type ApiPool interface {
	Get(symbol string) ApiGatewayService
}

type apiPool struct {
	apis map[string]ApiGatewayService
}

func NewApiPool(ctx context.Context, symbols []string, tradeRepository trade.TradeRepositoryAdmin, candleRepository GraphRepository) apiPool {
	apis := map[string]ApiGatewayService{}
	for _, symbol := range symbols {
		api := NewApiGatewayService(symbol, tradeRepository, candleRepository)
		go api.PullCandle(ctx)

		apis[symbol] = api
	}

	return apiPool{apis: apis}
}

func (p apiPool) Get(symbol string) ApiGatewayService {
	return p.apis[symbol]
}
