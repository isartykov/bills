package repository

import (
	"github.com/isartykov/user"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user.User) (int, error)
	GetUser(username, password string) (user.User, error)
}

type Bill interface {
	Create(userID int, bill user.Bill) (int, error)
	GetAll(userID int) ([]user.Bill, error)
	GetByID(userID, billID int) (user.Bill, error)
	Delete(userID, billID int) error
	Update(userID, billID int, input user.UpdateBillInput) error
	GetAllCalculate(userID int, input user.Times) ([]user.Bill, error)
}

type Repository struct {
	Authorization
	Bill
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Bill:          NewBillPostgres(db),
	}
}
