package gapi

import (
	"context"
	"time"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"github.com/gufir/money-management/val"
	"github.com/gufir/money-management/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {

	violations := ValidateForgotPasswordRequest(req)
	if violations != nil {
		return nil, InvalidArgumentError(violations)
	}

	user, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Email not found")
	}

	token := utils.RandomString(32)

	expires := time.Now().Add(15 * time.Minute)

	_, err = server.store.CreatePasswordReset(ctx, db.CreatePasswordResetParams{
		ID:        uuid.New(),
		UserID:    user.UserUuid,
		Token:     token,
		ExpiredAt: expires,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to set reset password token: %v", err)
	}

	taskPayload := &worker.PayloadSendForgotPassword{
		Email: user.Email,
		Token: token,
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.Queue(worker.QueueDefault),
	}

	err = server.taskDistributor.DistributeTaskResetPassword(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to enqueue reset password email task: %v", err)
	}

	return &pb.ForgotPasswordResponse{
		Message: "If the email exists, a password reset link has been sent.",
	}, nil
}

func ValidateForgotPasswordRequest(req *pb.ForgotPasswordRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	return
}

func ValidateResetPasswordRequest(req *pb.ResetPasswordRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidatePassword(req.GetNewPassword()); err != nil {
		violations = append(violations, fieldViolation("new_password", err))
	}
	return
}
