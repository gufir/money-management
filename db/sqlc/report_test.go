package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomreport(t *testing.T) {
	transaction := CreateRandomTransaction(t)

	arg := CreateReportUserParams{
		UserID: transaction.UserID,
		ID:     uuid.New(),
	}

	err := testQueries.CreateReportUser(context.Background(), arg)
	require.NoError(t, err)

}

func TestCreateReport(t *testing.T) {
	CreateRandomreport(t)
}
