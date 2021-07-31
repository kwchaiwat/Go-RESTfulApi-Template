package repository

import (
	model "go-restful-api-template/models"

	"gorm.io/gorm"
)

// ส่วนของ query จะทำที่นี่เท่านั้นไม่เกี่ยวกับ business logic
type customerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepositoryImpl(db *gorm.DB) customerRepositoryImpl {
	return customerRepositoryImpl{db: db}
}

func (r customerRepositoryImpl) GetAll() ([]model.Customer, error) {
	customers := []model.Customer{}
	tx := r.db.Find(&customers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return customers, nil
}

func (r customerRepositoryImpl) GetById(id int) (*model.Customer, error) {
	customer := model.Customer{}
	tx := r.db.Find(&customer, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &customer, nil
}
