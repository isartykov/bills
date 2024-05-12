package repository

import (
	"fmt"
	"strings"

	"github.com/isartykov/user"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type BillPostgres struct {
	db *sqlx.DB
}

func NewBillPostgres(db *sqlx.DB) *BillPostgres {
	return &BillPostgres{db: db}
}

func (r *BillPostgres) Create(userID int, bill user.Bill) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBillQuery := fmt.Sprintf("INSERT INTO %s (product, price, created_at) VALUES ($1, $2, $3) RETURNING id", billsTable)
	row := tx.QueryRow(createBillQuery, bill.Product, bill.Price, bill.CreatedAt)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersBillQuery := fmt.Sprintf("INSERT INTO %s (user_id, bill_id) VALUES ($1,$2)", usersBillsTable)
	_, err = tx.Exec(createUsersBillQuery, userID, id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *BillPostgres) GetAll(userID int) ([]user.Bill, error) {
	var bills []user.Bill

	query := fmt.Sprintf("SELECT bl.id, bl.product, bl.price, bl.created_at FROM %s bl INNER JOIN %s ub on bl.id = ub.bill_id WHERE ub.user_id = $1", billsTable, usersBillsTable)

	err := r.db.Select(&bills, query, userID)

	return bills, err
}

func (r *BillPostgres) GetByID(userID, billID int) (user.Bill, error) {
	var bill user.Bill

	query := fmt.Sprintf(`SELECT bl.id, bl.product, bl.price, bl.created_at FROM %s bl 
						INNER JOIN %s ub on bl.id = ub.bill_id WHERE ub.user_id = $1 AND ub.bill_id = $2`,
		billsTable, usersBillsTable)

	err := r.db.Get(&bill, query, userID, billID)

	return bill, err
}

func (r *BillPostgres) Delete(userID, billID int) error {
	query := fmt.Sprintf("DELETE FROM %s bl USING %s ub WHERE bl.id = ub.bill_id AND ub.user_id = $1 AND ub.bill_id = $2", billsTable, usersBillsTable)
	_, err := r.db.Exec(query, userID, billID)

	return err
}

func (r *BillPostgres) Update(userID, billID int, input user.UpdateBillInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Product != nil {
		setValues = append(setValues, fmt.Sprintf("product=$%d", argID))
		args = append(args, *input.Product)
		argID++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argID))
		args = append(args, *input.Price)
		argID++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s bl SET %s FROM %s ub WHERE bl.id=ub.bill_id AND ub.bill_id=$%d AND ub.user_id=$%d",
		billsTable, setQuery, usersBillsTable, argID, argID+1)

	args = append(args, billID, userID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *BillPostgres) GetAllCalculate(userID int, input user.Times) ([]user.Bill, error) {
	var bills []user.Bill

	query := fmt.Sprintf("SELECT bl.id, bl.product, bl.price, bl.created_at FROM %s bl INNER JOIN %s ub on bl.id = ub.bill_id WHERE ub.user_id = $1 AND bl.created_at BETWEEN $2 AND $3", billsTable, usersBillsTable)

	err := r.db.Select(&bills, query, userID, input.StartTime, input.EndTime)

	return bills, err
}
