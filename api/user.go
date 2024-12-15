package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/utils"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/labstack/echo/v4"
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
		CreatedAt: user.CreatedAt.Time,
		UserUuid:  user.UserUuid,
	}
}

func (server *Server) CreateUser(ctx echo.Context) error {
	var req CreateUserRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, ErrorResponse(err.Error()))
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		UserUuid:       uuid.New(),
	}

	user, err := server.store.CreateUser(ctx.Request().Context(), arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return ctx.JSON(http.StatusConflict, ErrorResponse("username or email already exists"))
		}

		return ctx.JSON(http.StatusInternalServerError, ErrorResponse("Failed to create User"))
	}

	res := NewUserResponse(user)

	return ctx.JSON(http.StatusCreated, res)

}
