package service

import (
	"github.com/isartykov/user"
	"github.com/isartykov/user/pkg/repository"
)

type Authorization interface {
	CreateUser(user user.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Bill interface {
	Create(userID int, bill user.Bill) (int, error)
	GetAll(userID int) ([]user.Bill, error)
	GetByID(userID, billID int) (user.Bill, error)
	Delete(userID, billID int) error
	Update(userID, billID int, input user.UpdateBillInput) error
	GetRandomBill(userID int) (user.Bill, error)
	GetAllCalculate(userID int, input user.Times) ([]user.Bill, error)
}

type Service struct {
	Authorization
	Bill
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Bill:          NewBillService(repos.Bill),
	}
}
