package service

import (
	"math/rand"
	"time"

	"github.com/isartykov/user"
	"github.com/isartykov/user/pkg/repository"
)

type BillService struct {
	repo repository.Bill
}

func NewBillService(repo repository.Bill) *BillService {
	return &BillService{repo: repo}
}

func (s *BillService) Create(userID int, bill user.Bill) (int, error) {
	bill.CreatedAt = time.Now().UTC()
	return s.repo.Create(userID, bill)
}

func (s *BillService) GetAll(userID int) ([]user.Bill, error) {
	return s.repo.GetAll(userID)
}

func (s *BillService) GetByID(userID, billID int) (user.Bill, error) {
	return s.repo.GetByID(userID, billID)
}

func (s *BillService) Delete(userID, billID int) error {
	return s.repo.Delete(userID, billID)
}

func (s *BillService) Update(userID, billID int, input user.UpdateBillInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userID, billID, input)
}

func (s *BillService) GetRandomBill(userID int) (user.Bill, error) {
	bills, err := s.repo.GetAll(userID)

	if err != nil {
		return user.Bill{}, err
	}

	t := rand.Int31n(int32(len(bills)))

	return bills[t], nil
}

func (s *BillService) GetAllCalculate(userID int, input user.Times) ([]user.Bill, error) {
	return s.repo.GetAllCalculate(userID, input)
}
