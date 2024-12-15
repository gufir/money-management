package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
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

// ExecTx executes a function within a database transaction.
// It creates a new `Queries` instance tied to the transaction.
func (store *SQLStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.Begin(ctx) // Begin a transaction
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}
	defer tx.Rollback(ctx) // Ensure rollback on error

	q := New(tx) // Create a `Queries` instance bound to this transaction
	if err := fn(q); err != nil {
		return err // Return the error from the callback
	}

	// Commit the transaction if the callback succeeds
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	return nil
}
