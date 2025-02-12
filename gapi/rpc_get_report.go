package gapi

import (
	"context"
	"time"

	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetReport(ctx context.Context, req *pb.GetReportByUserRequest) (*pb.GetReportByUserResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	if authPayload.Role != utils.AdminRole && authPayload.Role != utils.UserRole {
		return nil, status.Errorf(codes.PermissionDenied, "only admin and users can access reports")
	}

	report, err := server.store.GetReportByUser(ctx, authPayload.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get report: %v", err)
	}

	rsp := &pb.GetReportByUserResponse{
		UserId:       report.UserID.String(),
		TotalIncome:  float64(report.TotalIncome),
		TotalExpense: float64(report.TotalExpense),
	}

	return rsp, nil
}

func (server *Server) GetReportByDate(ctx context.Context, req *pb.GetReportByDateRequest) (*pb.GetReportByDateResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	if authPayload.Role != utils.AdminRole && authPayload.Role != utils.UserRole {
		return nil, status.Errorf(codes.PermissionDenied, "only admin and users can access reports")
	}

	layout := "2006-01-02"
	startDate, err := time.Parse(layout, req.StartDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid start date format: %v", err)
	}

	endDate, err := time.Parse(layout, req.EndDate)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid end date format: %v", err)
	}

	params := db.GetReportByDateParams{
		UserID:      authPayload.UserID,
		CreatedAt:   startDate,
		CreatedAt_2: endDate,
	}

	transactions, err := server.store.GetReportByDate(ctx, params)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get report by date: %v", err)
	}

	var pbReports []*pb.ReportByDate
	for _, tx := range transactions {
		pbReports = append(pbReports, &pb.ReportByDate{
			Id:          tx.ID.String(),
			UserId:      tx.UserID.String(),
			Amount:      float64(tx.Amount),
			Type:        tx.Type,
			Description: tx.Description,
			CreatedAt:   timestamppb.New(tx.CreatedAt),
		})
	}

	// Construct response
	rsp := &pb.GetReportByDateResponse{
		Reports: pbReports,
	}

	return rsp, nil
}
