package repository

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	OpeningDate time.Time
	AccountType string `gorm:"size:20"`
	Amount      float64
	Status      int
	Customer    Customer
	CustomerID  int
}

type AccountRepository interface {
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
}
