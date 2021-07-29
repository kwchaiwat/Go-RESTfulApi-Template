package repositories

import (
	"go-restful-api-template/models"

	"github.com/jmoiron/sqlx"
)

// ส่วนของ query จะทำที่นี่เท่านั้นไม่เกี่ยวกับ business logic
type bankRepositoryImpl struct {
	db *sqlx.DB
}

func NewBankRepositoryImpl(db *sqlx.DB) bankRepositoryImpl {
	return bankRepositoryImpl{db: db}
}

func (r bankRepositoryImpl) GetAll() ([]models.Bank, error) {
	banks := []models.Bank{}
	query := "SELECT * FROM banks"
	err := r.db.Select(&banks, query)
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func (r bankRepositoryImpl) GetById(id int) (*models.Bank, error) {
	bank := models.Bank{}
	query := "SELECT * FROM banks WHERE id=?"
	err := r.db.Get(&bank, query, id)
	if err != nil {
		return nil, err
	}
	return &bank, nil
}
