package gapi

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"github.com/gufir/money-management/utils"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateCategories(ctx context.Context, req *pb.UpdateCategoriesRequest) (*pb.UpdateCategoriesResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthorizedError(err)
	}

	if authPayload.Role != utils.AdminRole {
		return nil, status.Errorf(codes.PermissionDenied, "only admin can update categories")
	}

	CategoryID, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid category id: %v", err)
	}

	if _, err := server.store.GetCategoryById(ctx, CategoryID); err != nil {
		return nil, status.Errorf(codes.NotFound, "category not found: %v", err)
	}

	if existingCategory, err := server.store.GetCategoryByName(ctx, req.Name); err == nil && existingCategory.Name == req.Name {
		return nil, status.Errorf(codes.AlreadyExists, "category with name '%s' already exists", req.Name)
	} else if err != nil && err == sql.ErrNoRows {
		return nil, status.Errorf(codes.Internal, "failed to check if category exists: %v", err)
	}

	arg := db.UpdateCategoriesParams{
		ID: CategoryID,
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != "",
		},
	}

	category, err := server.store.UpdateCategories(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update category: %v", err)
	}

	rsp := &pb.UpdateCategoriesResponse{
		Category: ConvertCategory(category),
	}

	return rsp, nil

}
