package gapi

import (
	"context"

	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetCategories(ctx context.Context, req *pb.GetCategoriesRequest) (*pb.GetCategoriesResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	if authPayload.Role != utils.AdminRole && authPayload.Role != utils.UserRole {
		return nil, status.Errorf(codes.PermissionDenied, "only admin can list categories")
	}

	categories, err := server.store.GetAllCategories(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get categories: %v", err)
	}

	var pbCategories []*pb.Category
	for _, category := range categories {
		pbCategories = append(pbCategories, &pb.Category{
			Id:   category.ID.String(),
			Name: category.Name,
		})
	}

	rsp := &pb.GetCategoriesResponse{
		Categories: pbCategories,
	}

	return rsp, nil

}
