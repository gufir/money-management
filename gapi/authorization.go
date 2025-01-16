package gapi

import (
	"context"
	"fmt"
	"strings"

	"github.com/gufir/money-management/token"
	"google.golang.org/grpc/metadata"
)

const (
	AuthorizationHeader = "authorization"
	AuthorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context, accessibleRoles []string) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(AuthorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization token")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != AuthorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %s", authType)
	}

	accessToken := fields[1]
	payload, err := server.token.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %w", err)
	}

	if !hasPermission(payload.Role, accessibleRoles) {
		return nil, fmt.Errorf("permission denied")
	}

	return payload, nil
}

func hasPermission(userRole string, accessibleRoles []string) bool {
	for _, role := range accessibleRoles {
		if userRole == role {
			return true
		}
	}
	return false
}
