package repository

import "github.com/jmoiron/sqlx"

type accountRepositoryImpl struct {
	db *sqlx.DB
}

func NewAccountRepositoryImpl(db *sqlx.DB) accountRepositoryImpl {
	return accountRepositoryImpl{db: db}
}

func (r accountRepositoryImpl) Create(acc Account) (*Account, error) {
	query := "INSERT into accounts (customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	result, err := r.db.Exec(
		query,
		acc.CustomerID,
		acc.OpeningDate,
		acc.AccountType,
		acc.Amount,
		acc.Status,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	acc.AccountID = int(id)
	return &acc, nil
}

func (r accountRepositoryImpl) GetAll(customerID int) ([]Account, error) {
	query := "select * from accounts where customer_id=?"
	accounts := []Account{}
	err := r.db.Select(&accounts, query, customerID)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
