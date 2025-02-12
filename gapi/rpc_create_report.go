package gapi

import (
	"context"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateReportUser(ctx context.Context, req *pb.CreateReportRequest) (*pb.CreateReportResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid user id: %v", err)
	}

	if authPayload.UserID != userID {
		return nil, status.Errorf(codes.PermissionDenied, "can't create transaction for another user")
	}

	arg := db.CreateReportUserParams{
		ID:     uuid.New(),
		UserID: authPayload.UserID,
	}

	report, err := server.store.CreateReportUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create report: %v", err)
	}

	rsp := &pb.CreateReportResponse{
		Report: ConverReport(report),
	}

	return rsp, nil
}
