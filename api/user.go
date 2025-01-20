package api

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/utils"
	"github.com/jackc/pgx/v5/pgconn"
)

type CreateUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
	FullName string `json:"full_name" binding:"required"`
}

type CreateUserResponse struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UserUuid  uuid.UUID `json:"user_uuid"`
}

func NewUserResponse(user db.User) CreateUserResponse {
	return CreateUserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UserUuid:  user.UserUuid,
	}
}

func (server *Server) CreateUser(ctx *gin.Context) {
	var req CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		UserUuid:       uuid.New(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			ctx.JSON(http.StatusConflict, errorResponse("username or email already exists"))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse("Failed to create User"))
		return
	}

	res := NewUserResponse(user)

	ctx.JSON(http.StatusOK, res)

}

type LoginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	SessionID             uuid.UUID          `json:"id"`
	AccessToken           string             `json:"access_token"`
	AccressTokenExpiredAt time.Time          `json:"access_token_expired_at"`
	RefreshToken          string             `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time          `json:"refresh_token_expired_at"`
	User                  CreateUserResponse `json:"user"`
}

func (server *Server) LoginUser(ctx *gin.Context) {
	var req LoginRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
	}

	user, err := server.store.GetUserByUsername(ctx, req.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	err = utils.CheckPasswordHash(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err.Error()))
		return
	}

	accessToken, accessPayload, err := server.token.CreateToken(
		user.UserUuid,
		user.Username,
		user.Role,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	refreshToken, refreshPayload, err := server.token.CreateToken(
		user.UserUuid,
		user.Username,
		user.Role,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		UserID:       refreshPayload.UserID,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	resp := LoginResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccressTokenExpiredAt: accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiredAt: accessPayload.ExpiredAt,
		User:                  NewUserResponse(user),
	}

	ctx.JSON(http.StatusOK, resp)
}
