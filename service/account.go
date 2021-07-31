package service

import (
	"time"
)

type NewAccountRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type UpdateAccountRequest struct {
	CustomerID  int     `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      int     `json:"status"`
}

type AccountResponse struct {
	Customer    CustomerResponse `json:"customer"`
	ID          uint             `json:"id"`
	OpeningDate time.Time        `json:"opening_date"`
	AccountType string           `json:"account_type"`
	Amount      float64          `json:"amount"`
	Status      int              `json:"status"`
}

type AccountService interface {
	NewAccount(int, NewAccountRequest) (*AccountResponse, error)
	GetAccounts(id int) ([]AccountResponse, error)
	UpdateAccount(int, UpdateAccountRequest) (*AccountResponse, error)
}
