package repositories

import (
	"github.com/jmoiron/sqlx"
)

// ส่วนของ query จะทำที่นี่เท่านั้นไม่เกี่ยวกับ business logic
type customerRepositoryImpl struct {
	db *sqlx.DB
}

func NewCustomerRepositoryImpl(db *sqlx.DB) customerRepositoryImpl {
	return customerRepositoryImpl{db: db}
}

func (r customerRepositoryImpl) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "SELECT * FROM customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryImpl) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "SELECT * FROM customers WHERE id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}
