package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lovelyoyrmia/ecommerce/pkg/config"
)

type Database struct {
	DB *pgxpool.Pool
}

func NewDatabase(ctx context.Context, conf config.Config) *Database {
	sqlDriver, err := pgxpool.New(ctx, conf.DBUrl)
	if err != nil {
		return nil
	}
	return &Database{
		DB: sqlDriver,
	}
}
