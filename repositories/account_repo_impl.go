package repositories

import "github.com/jmoiron/sqlx"

type AccountRepositoryImpl struct {
	db *sqlx.DB
}

func (r AccountRepository) Create(acc Account) (*Account, error) {
	return nil, nil
}

func (r AccountRepository) GetAll(customerID int) ([]Account, error) {
	return nil, nil
}
