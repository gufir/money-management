package gapi

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/token"
	"github.com/gufir/money-management/utils"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymmetricKey:   utils.RandomString(32),
		AccessTokenDuration: time.Hour,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func NewContextWithBearerToken(t *testing.T, token token.Maker, UserUuid uuid.UUID, username string, role string, duration time.Duration) context.Context {
	accessToken, _, err := token.CreateToken(UserUuid, username, role, duration)
	require.NoError(t, err)

	bearerToken := fmt.Sprintf("%s %s", AuthorizationBearer, accessToken)
	md := metadata.MD{
		AuthorizationHeader: []string{
			bearerToken,
		},
	}

	return metadata.NewIncomingContext(context.Background(), md)
}
