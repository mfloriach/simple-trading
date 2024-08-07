package utils

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var TimeLineDb *pgxpool.Pool

func GetTimeLineDb(ctx context.Context) error {
	if TimeLineDb != nil {
		return nil
	}

	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	conn, err := pgxpool.New(ctx, "postgres://"+user+":"+password+"@localhost:5432/"+db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	TimeLineDb = conn

	return nil
}
