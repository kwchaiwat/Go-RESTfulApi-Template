package repository

import (
	"database/sql"
	"go-restful-api-template/logs"
	model "go-restful-api-template/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type accountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepositoryImpl(db *gorm.DB) accountRepositoryImpl {
	return accountRepositoryImpl{db: db}
}

func (r accountRepositoryImpl) Create(acc model.Account) (*model.Account, error) {
	tx := r.db.Create(&acc)
	if tx.Error != nil {
		logs.Error(tx.Error)
		return nil, tx.Error
	}
	return &acc, nil
}

func (r accountRepositoryImpl) GetAll(customerID int) ([]model.Account, error) {
	account := []model.Account{}
	tx := r.db.Preload(clause.Associations).Where("customer_id", customerID).Find(&account)
	if tx.Error != nil {
		logs.Error(tx.Error)
		return nil, tx.Error
	}
	return account, nil
}

func (r accountRepositoryImpl) GetById(id int) (*model.Account, error) {
	acc := model.Account{}
	tx := r.db.Preload(clause.Associations).First(&acc, id)
	if tx.Error != nil {
		logs.Error(tx.Error)
		return nil, tx.Error
	}
	return &acc, nil
}

func (r accountRepositoryImpl) Update(accountID int, acc model.Account) (*model.Account, error) {
	tx := r.db.Model(&model.Account{}).First(&model.Account{}, accountID).Where("id=@accountID", sql.Named("accountID", accountID)).Updates(acc)
	if tx.Error != nil {
		logs.Error(tx.Error)
		return nil, tx.Error
	}
	return &acc, nil
}

func (r accountRepositoryImpl) Delete(accountID int) error {
	acc := model.Account{}
	tx := r.db.First(&acc, accountID).Delete(&acc, accountID)
	if tx.Error != nil {
		logs.Error(tx.Error)
		return tx.Error
	}
	return nil
}
