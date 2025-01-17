package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/gufir/money-management/db/sqlc"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var req renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	refreshPayload, err := server.token.VerifyToken(req.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err.Error()))
		return
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusUnauthorized, errorResponse("refresh token not found"))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
	}

	if session.UserID != refreshPayload.UserID {
		ctx.JSON(http.StatusUnauthorized, errorResponse("incorrect sessions user"))
		return
	}

	if session.RefreshToken != req.RefreshToken {
		ctx.JSON(http.StatusUnauthorized, errorResponse("Mismatched refresh token"))
		return
	}

	if time.Now().After(session.ExpiresAt) {
		ctx.JSON(http.StatusUnauthorized, errorResponse("refresh token is expired"))
		return
	}

	accessToken, accessPayload, err := server.token.CreateToken(
		refreshPayload.UserID,
		refreshPayload.Username,
		refreshPayload.Role,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	resp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, resp)

}
