package infrastructure

import (
	"context"

	"github.com/wit-switch/assessment-tax/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresClient(ctx context.Context, cfg *config.PostgresConfig) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, cfg.URL)
	if err != nil {
		return nil, err
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
