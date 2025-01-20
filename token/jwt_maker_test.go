package token

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gufir/money-management/utils"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(utils.RandomString(32))
	require.NoError(t, err)

	username := utils.RandomName()
	userID := uuid.New()
	role := utils.UserRole
	duration := time.Hour

	IssuedAt := time.Now()
	ExpiredAt := IssuedAt.Add(duration)

	token, payload, err := maker.CreateToken(userID, role, username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.UserID)
	require.Equal(t, role, payload.Role)
	require.WithinDuration(t, IssuedAt, payload.CreatedAt, time.Second)
	require.WithinDuration(t, ExpiredAt, payload.ExpiredAt, time.Second)

}
