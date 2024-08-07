package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"example.com/grpc-todo/utils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()

	utils.GetTimeLineDb(ctx)

	createTables(ctx)
	createTradesTables(ctx)

	// viewsPerSecond(ctx)
	viewsPerDay(ctx)
	viewsPerMinute(ctx)
	viewsPerFiveMinute(ctx)
}

func createTradesTables(ctx context.Context) {
	if _, err := utils.TimeLineDb.Exec(ctx, `
	CREATE TABLE trades (
		id SERIAL PRIMARY KEY, 
		user_id INT NOT NULL,
		symbol VARCHAR(50) NOT NULL, 
		action INT NOT NULL,
		amount INT NOT NULL
	);`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}
}

func createTables(ctx context.Context) {
	if _, err := utils.TimeLineDb.Exec(ctx, `
	CREATE TABLE stock_price (
		time TIMESTAMPTZ NOT NULL,
		symbol VARCHAR(10) NOT NULL,
		price INT NOT NULL,
		PRIMARY KEY (time, symbol)
	);

	SELECT create_hypertable('stock_price', by_range('time'));

	CREATE INDEX ix_symbol_time ON stock_price (symbol, time DESC);

	SELECT add_retention_policy('stock_price', INTERVAL '10 years');
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}
}

func viewsPerSecond(ctx context.Context) {
	if _, err := utils.TimeLineDb.Exec(ctx, `
	CREATE MATERIALIZED VIEW one_second_candle
	WITH (timescaledb.continuous) AS
		SELECT
			time_bucket('1 second', time) AS time,
			symbol,
			FIRST(price, time) AS "open",
			MAX(price) AS high,
			MIN(price) AS low,
			LAST(price, time) AS "close"
		FROM stock_price
		GROUP BY time, symbol;
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}

	if _, err := utils.TimeLineDb.Exec(ctx, `
	SELECT add_continuous_aggregate_policy('one_second_candle',
    start_offset => INTERVAL '3 days',
    end_offset => INTERVAL '1 day',
    schedule_interval => INTERVAL '1 day');
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}
}

func viewsPerDay(ctx context.Context) {
	if _, err := utils.TimeLineDb.Exec(ctx, `
	CREATE MATERIALIZED VIEW one_day_candle
	WITH (timescaledb.continuous) AS
		SELECT
			time_bucket('1 day', time) AS bucket,
			symbol,
			FIRST(price, time) AS "open",
			MAX(price) AS high,
			MIN(price) AS low,
			LAST(price, time) AS "close"
		FROM stock_price
		GROUP BY bucket, symbol;
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}

	if _, err := utils.TimeLineDb.Exec(ctx, `
	SELECT add_continuous_aggregate_policy('one_day_candle',
    start_offset => INTERVAL '3 days',
    end_offset => INTERVAL '1 day',
    schedule_interval => INTERVAL '1 day');
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}
}

func viewsPerMinute(ctx context.Context) {
	if _, err := utils.TimeLineDb.Exec(ctx, `
	CREATE MATERIALIZED VIEW one_minute_candle
	WITH (timescaledb.continuous) AS
		SELECT
			time_bucket('1 minute', time) AS bucket,
			symbol,
			FIRST(price, time) AS "open",
			MAX(price) AS high,
			MIN(price) AS low,
			LAST(price, time) AS "close"
		FROM stock_price
		GROUP BY bucket, symbol;
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}

	if _, err := utils.TimeLineDb.Exec(ctx, `
	SELECT add_continuous_aggregate_policy('one_minute_candle',
    start_offset => INTERVAL '3 minute',
    end_offset => INTERVAL '1 minute',
    schedule_interval => INTERVAL '1 minute');
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}
}

func viewsPerFiveMinute(ctx context.Context) {
	if _, err := utils.TimeLineDb.Exec(ctx, `
	CREATE MATERIALIZED VIEW five_minute_candle
	WITH (timescaledb.continuous) AS
		SELECT
			time_bucket('5 minute', time) AS bucket,
			symbol,
			FIRST(price, time) AS "open",
			MAX(price) AS high,
			MIN(price) AS low,
			LAST(price, time) AS "close"
		FROM stock_price
		GROUP BY bucket, symbol;
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}

	if _, err := utils.TimeLineDb.Exec(ctx, `
	SELECT add_continuous_aggregate_policy('five_minute_candle',
    start_offset => INTERVAL '15 minute',
    end_offset => INTERVAL '5 minute',
    schedule_interval => INTERVAL '15 minute');
	`); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to insert data into database: %v\n", err)
	}
}
