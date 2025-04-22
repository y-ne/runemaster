package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(connString string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, fmt.Errorf("unable to parse connection string: %w", err)
	}

	config.MaxConns = int32(10)
	config.MinConns = int32(2)
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = time.Minute

	ctx := context.Background()
	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("unable to ping database: %w", err)
	}

	return pool, nil
}
