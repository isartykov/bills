package user

import (
	"errors"
	"time"
)

type Bill struct {
	Id        int       `json:"id" db:"id"`
	Product   string    `json:"product" db:"product"`
	Price     float64   `json:"price" db:"price"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UsersBill struct {
	Id     int
	UserId int
	ListId int
}

type UpdateBillInput struct {
	Product *string  `json:"product"`
	Price   *float64 `json:"price"`
}

func (i UpdateBillInput) Validate() error {
	if i.Product == nil && i.Price == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type Times struct {
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}
