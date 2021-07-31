package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type accountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepositoryImpl(db *gorm.DB) accountRepositoryImpl {
	return accountRepositoryImpl{db: db}
}

func (r accountRepositoryImpl) Create(acc Account) (*Account, error) {
	tx := r.db.Create(&acc)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return nil, tx.Error
	}
	fmt.Println(acc)
	return &acc, nil
}

func (r accountRepositoryImpl) GetAll(customerID int) ([]Account, error) {
	account := []Account{}
	tx := r.db.Where("customer_id", customerID).Find(&account)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return account, nil
}
