package gapi

import (
	"context"
	"time"

	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.ResetPasswordResponse, error) {

	violations := ValidateResetPasswordRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	tokenObj, err := server.store.GetPasswordResetToken(ctx, req.GetToken())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Invalid or expired reset token")
	}

	if tokenObj.IsUsed || time.Now().After(tokenObj.ExpiredAt) {
		return nil, status.Errorf(codes.NotFound, "Reset token has expired or already used")
	}

	hashedPassword, err := utils.HashPassword(req.GetNewPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to hash new password: %v", err)
	}

	_, err = server.store.UpdateUserPassword(ctx, db.UpdateUserPasswordParams{
		UserUuid:       tokenObj.UserID,
		HashedPassword: hashedPassword,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update password: %v", err)
	}

	err = server.store.UsePasswordResetToken(ctx, tokenObj.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to mark reset token as used: %v", err)
	}

	return &pb.ResetPasswordResponse{
		Message: "Password has been reset successfully",
	}, nil
}
