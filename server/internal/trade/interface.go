package trade

import (
	"context"
	"time"
)

type Action int8

const (
	ActionPut Action = iota
	ActionCall
	ActionTakeProfit
	ActionStopLose
)

type TradeAction struct {
	Symbol string
	Amount int64
}

type Trade struct {
	ID        string `bson:"_id,omitempty"`
	UserID    int
	Symbol    string
	Amount    int64
	Action    Action
	CreatedAt time.Time
}

type TradeService interface {
	Put(context.Context, int, TradeAction) error
	Call(context.Context, int, TradeAction) error
	TakeProfit(context.Context, int, TradeAction) error
	StopLose(context.Context, int, TradeAction) error
}

type TradeRepository interface {
	Set(context.Context, Trade) error
	GetByUserID(context.Context, int, Action) ([]Trade, error)
}

type TradeRepositoryAdmin interface {
	Get(context.Context, string, Action) ([]Trade, error)
	Delete(context.Context, string) error
}
