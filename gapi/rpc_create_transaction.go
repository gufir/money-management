package gapi

import (
	"context"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.CreateTransactionResponse, error) {

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %v", err)
	}

	if authPayload.UserID != userID {
		return nil, status.Errorf(codes.PermissionDenied, "can't create transaction for another user")
	}

	if !validateTransactionType(req.GetType()) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid transaction type: %v", req.Type)
	}

	var categoryID uuid.UUID
	if req.GetCategoryId() != "" {
		categoryID = uuid.MustParse(req.GetCategoryId())
	} else if req.GetCategoryName() != "" {
		category, err := server.store.GetCategoryByName(ctx, req.GetCategoryName())
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "category not found: %v", err)
		}
		categoryID = category.ID
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "category_id or category_name is required")
	}

	arg := db.CreateTransactionParams{
		ID:          uuid.New(),
		UserID:      userID,
		Amount:      req.Amount,
		Type:        req.Type,
		Description: req.Description,
		CategoryID:  categoryID,
	}

	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create transaction: %v", err)
	}

	rsp := &pb.CreateTransactionResponse{
		Transaction: ConvertTransaction(transaction),
	}

	return rsp, nil
}

func validateTransactionType(transactionType string) bool {
	validTypes := []string{"income", "expense"}
	for _, t := range validTypes {
		if t == transactionType {
			return true
		}
	}
	return false
}
