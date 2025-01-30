package db

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/gufir/money-management/utils"
	"github.com/stretchr/testify/require"
)

func CreateRandomCategories(t *testing.T) Category {
	arg := CreateCategoriesParams{
		ID:   uuid.New(),
		Name: utils.RandomString(6),
	}

	categories, err := testQueries.CreateCategories(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, arg.ID, categories.ID)
	require.Equal(t, arg.Name, categories.Name)

	return categories
}

func TestCategories(t *testing.T) {
	CreateRandomCategories(t)
}
