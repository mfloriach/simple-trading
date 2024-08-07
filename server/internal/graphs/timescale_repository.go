package graphs

import (
	"context"
	"fmt"
	"time"

	"example.com/grpc-todo/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TimeRange int8

const (
	TimeRangeSecond TimeRange = iota
	TimeRangeMinute
	TimeRangeFiveMinute
	TimeRangeDay
)

type GraphRepository interface {
	Insert(ctx context.Context, symbol string, price int) error
	GetBySymbol(ctx context.Context, symbol string, timeRange TimeRange) ([]Candle, error)
}

type TimescaleRepo struct {
	db *pgxpool.Pool
}

func NewTimescaleRepo() GraphRepository {
	return &TimescaleRepo{
		db: utils.TimeLineDb,
	}
}

func (r *TimescaleRepo) Insert(ctx context.Context, symbol string, price int) error {
	queryInsertMetadata := `INSERT INTO stock_price (time, symbol, price) VALUES ($1, $2, $3);`
	_, err := r.db.Exec(ctx, queryInsertMetadata, time.Now(), symbol, price)

	return err
}

func (r *TimescaleRepo) GetBySymbol(ctx context.Context, symbol string, timeRange TimeRange) ([]Candle, error) {
	var (
		rows pgx.Rows
		err  error
	)

	switch timeRange {
	case TimeRangeSecond:
		rows, err = r.getBySecond(ctx, symbol)
	case TimeRangeMinute:
		rows, err = r.getByMinute(ctx, symbol)
	case TimeRangeFiveMinute:
		rows, err = r.getByFiveMinute(ctx, symbol)
	case TimeRangeDay:
		rows, err = r.getByFDay(ctx, symbol)
	default:
		return []Candle{}, fmt.Errorf("time range not exist: %v", err)
	}

	var results []Candle
	for rows.Next() {
		var r Candle
		if err = rows.Scan(&r.Time, &r.Symbol, &r.High, &r.Open, &r.Close, &r.Low); err != nil {
			return []Candle{}, err
		}
		results = append(results, r)
	}

	if rows.Err() != nil {
		return []Candle{}, err
	}

	return results, nil
}

func (r *TimescaleRepo) getBySecond(ctx context.Context, symbol string) (pgx.Rows, error) {
	return r.db.Query(ctx, `
		SELECT * FROM one_second_candle
		WHERE symbol = $1 AND bucket >= NOW() - INTERVAL '14 days'
		ORDER BY time DESC;
		`, symbol)
}

func (r *TimescaleRepo) getByMinute(ctx context.Context, symbol string) (pgx.Rows, error) {
	return r.db.Query(ctx, `
		SELECT * FROM one_minute_candle
		WHERE symbol = $1 AND bucket >= NOW() - INTERVAL '14 days'
		ORDER BY bucket DESC;
		`, symbol)
}

func (r *TimescaleRepo) getByFiveMinute(ctx context.Context, symbol string) (pgx.Rows, error) {
	return r.db.Query(ctx, `
		SELECT * FROM five_minute_candle
		WHERE symbol = $1 AND bucket >= NOW() - INTERVAL '14 days'
		ORDER BY bucket DESC;
		`, symbol)
}

func (r *TimescaleRepo) getByFDay(ctx context.Context, symbol string) (pgx.Rows, error) {
	return r.db.Query(ctx, `
		SELECT * FROM one_day_candle
		WHERE symbol = $1 AND bucket >= NOW() - INTERVAL '14 days'
		ORDER BY bucket DESC;
		`, symbol)
}
