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

	parseUUID, err := uuid.Parse(req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %v", err)
	}

	if authPayload.UserID != parseUUID {
		return nil, status.Errorf(codes.PermissionDenied, "can't create transaction for another user")
	}

	arg := db.CreateTransactionParams{
		ID:          uuid.New(),
		UserID:      parseUUID,
		Amount:      req.Amount,
		Type:        req.Type,
		Description: req.Description,
	}

	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create transaction: %v", err)
	}

	rsp := &pb.CreateTransactionResponse{
		Transaction: ConvertTransaction(transaction),
	}

	// transaction := &pb.Transaction{
	// 	Amount:     req.Amount,
	// 	Type:       req.Type,
	// 	Descrition: req.Description,
	// 	CreatedAt:  timestamppb.New(time.Now()),
	// 	UpdatedAt:  timestamppb.New(time.Now()),
	// }

	return rsp, nil
}
