package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

// Store provides all database functions, including transactions.
type SQLStore struct {
	db *pgxpool.Pool
	*Queries
}

// NewStore initializes a new Store instance with a pgx connection pool.
func NewStore(db *pgxpool.Pool) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
