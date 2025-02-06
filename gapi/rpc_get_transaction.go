package gapi

import (
	"context"

	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetTransaction(ctx context.Context, req *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	if authPayload.Role != utils.AdminRole && authPayload.Role != utils.UserRole {
		return nil, status.Errorf(codes.PermissionDenied, "only admin can list categories")
	}

	transaction, err := server.store.GetTransactionByuserId(ctx, authPayload.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get transaction: %v", err)
	}

	var pbTx []*pb.Transaction
	for _, tx := range transaction {
		pbTx = append(pbTx, &pb.Transaction{
			Id:          tx.ID.String(),
			Amount:      tx.Amount,
			Type:        tx.Type,
			Description: tx.Description,
			CategoryId:  tx.CategoryID.String(),
			UserId:      tx.UserID.String(),
		})
	}

	rsp := &pb.GetTransactionResponse{
		Transaction: pbTx,
	}

	return rsp, nil

}
