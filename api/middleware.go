package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gufir/money-management/token"
)

const (
	authorizationHeaderkey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(token token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeaderkey := ctx.GetHeader(authorizationHeaderkey)
		if len(authorizationHeaderkey) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		fields := strings.Fields(authorizationHeaderkey)
		if len(fields) < 2 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("authorization type is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		accessToken := fields[1]
		payload, err := token.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err.Error()))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()

	}
}
