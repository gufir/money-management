package gapi

import (
	db "github.com/gufir/money-management/db/sqlc"
	"github.com/gufir/money-management/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertUser(user db.User) *pb.User {
	return &pb.User{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UserUuid:  user.UserUuid.String(),
	}
}

func ConvertTransaction(transaction db.Transaction) *pb.Transaction {
	return &pb.Transaction{
		Amount:      transaction.Amount,
		Type:        transaction.Type,
		Description: transaction.Description,
		CreatedAt:   timestamppb.New(transaction.CreatedAt),
		UpdatedAt:   timestamppb.New(transaction.UpdatedAt),
		UserId:      transaction.UserID.String(),
	}
}
