package gapi

import (
	"context"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.UpdateTransactionResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	violations := ValidateUpdateTransactionRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	userID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %v", err)
	}

	if authPayload.UserID != userID {
		return nil, status.Errorf(codes.PermissionDenied, "can't create transaction for another user")
	}

	TransactionID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid transaction id: %v", err)
	}

	arg := db.UpdateTransactionParams{
		ID:     TransactionID,
		UserID: userID,
		Amount: pgtype.Int8{
			Int64: req.GetAmount(),
			Valid: req.Amount != nil,
		},
		Type: req.Type,
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
	}

	if req.GetCategoryId() != "" {
		arg.CategoryID = uuid.MustParse(req.GetCategoryId())
	}

	transaction, err := server.store.UpdateTransaction(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update transaction: %v", err)
	}

	rsp := &pb.UpdateTransactionResponse{
		Transaction: ConvertTransaction(transaction),
	}

	return rsp, nil

}

func ValidateUpdateTransactionRequest(req *pb.UpdateTransactionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetAmount() < 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "amount",
			Description: "amount can't be negative",
		})
	}

	if req.GetDescription() == "" {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "description",
			Description: "description can't be empty",
		})
	}

	if !validateTransactionType(req.GetType()) {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field:       "type",
			Description: "invalid transaction type",
		})

		return violations
	}

	return violations
}
