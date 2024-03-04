package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type KeyUser struct{}

type Store interface {
	Querier
	ExecTx(ctx context.Context, fn func(*Queries) error) error
	CreateUserTx(ctx context.Context, userParams CreateUserParams) (User, error)
	CreateCartUserTx(ctx context.Context, req CreateCartTx) error
}

type SQLStore struct {
	*Queries
	ConnPool *pgxpool.Pool
}

func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		Queries:  New(connPool),
		ConnPool: connPool,
	}
}

func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.ConnPool.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx error: %d", rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
