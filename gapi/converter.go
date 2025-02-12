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

func ConvertCategory(category db.Category) *pb.Category {
	return &pb.Category{
		Id:   category.ID.String(),
		Name: category.Name,
	}
}

func ConverReport(report db.Report) *pb.Report {
	return &pb.Report{
		Id:           report.ID.String(),
		UserId:       report.UserID.String(),
		Period:       report.Period,
		TotalIncome:  float64(report.TotalIncome.Exp),
		TotalExpense: float64(report.TotalExpense.Exp),
	}
}
