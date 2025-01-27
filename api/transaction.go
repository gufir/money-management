package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/token"
)

type CreateTransactionRequest struct {
	Amount      int64  `json:"amount" binding:"required,min=1"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"required"`
}

type CreateTransactionResponse struct {
	Amount      int64     `json:"amount"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	UserID      uuid.UUID `json:"user_uuid"`
	CreatedAt   string    `json:"created_at"`
}

func NewTransactionResponse(transaction db.Transaction) CreateTransactionResponse {
	return CreateTransactionResponse{
		Amount:      transaction.Amount,
		Description: transaction.Description,
		Type:        transaction.Type,
		UserID:      transaction.UserID,
	}
}

func (server *Server) CreateUserTransaction(ctx *gin.Context) {

	authPayload, exists := ctx.Get(authorizationPayloadKey)
	if !exists {
		ctx.JSON(http.StatusUnauthorized, errorResponse("authorization payload is missing"))
		return
	}

	payload, ok := authPayload.(*token.Payload)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, errorResponse("invalid authorization payload"))
		return
	}

	var req CreateTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	arg := db.CreateTransactionParams{
		Amount:      req.Amount,
		Description: req.Description,
		Type:        req.Type,
		UserID:      payload.UserID,
	}

	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	response := NewTransactionResponse(transaction)
	ctx.JSON(http.StatusOK, response)
}
