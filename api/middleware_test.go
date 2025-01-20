package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gufir/money-management/token"
	"github.com/gufir/money-management/utils"
	"github.com/stretchr/testify/require"
)

func AddAuthorization(
	t *testing.T,
	request *http.Request,
	token token.Maker,
	authorizationType string,
	userID uuid.UUID,
	username string,
	role string,
	duration time.Duration,
) {
	accessToken, payload, err := token.CreateToken(userID, username, role, duration)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	authorizationHeader := fmt.Sprintf("%s %s", authorizationType, accessToken)
	request.Header.Set(authorizationHeaderkey, authorizationHeader)
}

func TestAuthMiddleware(t *testing.T) {
	userID := uuid.New()
	role := utils.UserRole
	username := utils.RandomName()

	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, token token.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, authorizationTypeBearer, userID, username, role, time.Hour)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},

		{
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},

		{
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, "unsupported", userID, username, role, time.Minute)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},

		{
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, "", userID, username, role, time.Minute)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},

		{
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request, token token.Maker) {
				AddAuthorization(t, request, token, authorizationTypeBearer, userID, username, role, -time.Minute)
			},

			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			server := NewTestServer(t, nil)
			authPath := "/auth"
			server.router.GET(
				authPath,
				authMiddleware(server.token),
				func(ctx *gin.Context) {
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(http.MethodGet, authPath, nil)
			require.NoError(t, err)

			tc.setupAuth(t, request, server.token)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
