package db

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gufir/money-management/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func CreateRandomTransaction(t *testing.T) Transaction {
	user := CreateRandomUser(t)
	categories := CreateRandomCategories(t)

	arg := CreateTransactionParams{
		ID:          uuid.New(),
		UserID:      user.UserUuid,
		Amount:      utils.RandomInt(10, 1000),
		Type:        utils.Expense,
		CategoryID:  categories.ID,
		Description: pgtype.Text{String: utils.RandomString(10), Valid: true},
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)
	require.Equal(t, transaction.UserID, arg.UserID)
	require.Equal(t, transaction.Amount, arg.Amount)
	require.Equal(t, transaction.Type, arg.Type)
	require.Equal(t, transaction.CategoryID, arg.CategoryID)

	require.NotEmpty(t, transaction.CreatedAt)
	require.NotEmpty(t, transaction.UpdatedAt)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	CreateRandomTransaction(t)
}

func TestGetTransaction(t *testing.T) {
	transaction1 := CreateRandomTransaction(t)
	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.ID)

	require.NoError(t, err)
	require.Equal(t, transaction1.UserID, transaction2.UserID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.Type, transaction2.Type)

	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)

}

func TestGetTransactionbyType(t *testing.T) {
	transaction1 := CreateRandomTransaction(t)
	params := GetTransactionByTypeParams{
		Type:   transaction1.Type,
		UserID: transaction1.UserID,
	}

	transactions, err := testQueries.GetTransactionByType(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, transactions)

	transaction2 := transactions[0]

	require.Equal(t, transaction1.UserID, transaction2.UserID)
	require.Equal(t, transaction1.Amount, transaction2.Amount)
	require.Equal(t, transaction1.Type, transaction2.Type)
	require.WithinDuration(t, transaction1.CreatedAt, transaction2.CreatedAt, time.Second)
}
