package gapi

import (
	"context"
	"errors"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}

		return nil, status.Errorf(codes.Internal, "failed to find user")
	}

	err = utils.CheckPasswordHash(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "incorrect password: %v", err)
	}

	accessToken, accessPayload, err := server.token.CreateToken(
		user.UserUuid,
		user.Username,
		user.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %v", err)
	}

	refreshToken, refreshPayload, err := server.token.CreateToken(
		user.UserUuid,
		user.Username,
		user.Role,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token: %v", err)
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           uuid.UUID(refreshPayload.ID),
		UserID:       user.UserUuid,
		RefreshToken: refreshToken,
		ExpiresAt:    refreshPayload.ExpiredAt,
		CreatedAt:    refreshPayload.CreatedAt,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh session: %v", err)
	}

	rsp := &pb.LoginUserResponse{
		User:                  ConvertUser(user),
		Id:                    session.ID.String(),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiredAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiredAt: timestamppb.New(refreshPayload.ExpiredAt),
	}

	return rsp, nil
}
