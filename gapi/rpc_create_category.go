package gapi

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCategories(ctx context.Context, req *pb.CreateCategoriesRequest) (*pb.CreateCategoriesResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	if authPayload.Role != utils.AdminRole {
		return nil, status.Errorf(codes.PermissionDenied, "only admin can create categories")
	}

	existingCategory, err := server.store.GetCategoryByName(ctx, req.Name)
	if err != nil && err != sql.ErrNoRows {
		return nil, status.Errorf(codes.Internal, "failed to check if category exists: %v", err)
	}

	if existingCategory.Name != "" && existingCategory.Name == req.Name {
		return nil, status.Errorf(codes.AlreadyExists, "category with name '%s' already exists", req.Name)
	}

	arg := db.CreateCategoriesParams{
		ID:   uuid.New(),
		Name: req.Name,
	}

	category, err := server.store.CreateCategories(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create category: %v", err)
	}

	rsp := &pb.CreateCategoriesResponse{
		Category: ConvertCategory(category),
	}

	return rsp, nil
}
