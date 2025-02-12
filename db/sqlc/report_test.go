package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func CreateRandomreport(t *testing.T) Report {
	transaction := CreateRandomTransaction(t)

	arg := CreateReportUserParams{
		UserID: transaction.UserID,
		ID:     uuid.New(),
	}

	report, err := testQueries.CreateReportUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, report)
	require.Equal(t, transaction.UserID, report.UserID)
	return report

}

func CreateRandomMonthlyReport(t *testing.T) Report {
	transaction := CreateRandomTransaction(t)

	report, err := testQueries.CreateMonthlyReport(context.Background(), CreateMonthlyReportParams{
		UserID: transaction.UserID,
		ID:     uuid.New(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, report)
	require.Equal(t, transaction.UserID, report.UserID)
	return report
}

func TestCreateReport(t *testing.T) {
	CreateRandomreport(t)
}

func TestCreateMonthlyReport(t *testing.T) {
	CreateRandomMonthlyReport(t)
}
