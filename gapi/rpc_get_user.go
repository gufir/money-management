package gapi

import (
	"context"

	"github.com/gufir/money-management/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	user, err := server.store.GetUserByUserId(ctx, authPayload.UserID)
	if err != nil {
		return nil, err
	}

	var pbUser []*pb.User
	pbUser = append(pbUser, &pb.User{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.FullName,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
		UserUuid:  user.UserUuid.String(),
	})

	rsp := &pb.GetUserResponse{
		User: pbUser,
	}

	return rsp, nil
}
