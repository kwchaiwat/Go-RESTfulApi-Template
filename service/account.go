package service

import (
	"database/sql"
	"go-restful-api-template/errs"
	"go-restful-api-template/logs"
	model "go-restful-api-template/models"
	"go-restful-api-template/repository"
	"strings"
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
	GetAccount(id int) ([]AccountResponse, error)
	UpdateAccount(int, UpdateAccountRequest) (*AccountResponse, error)
	DeleteAccount(id int) error
}

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {

	// Validate
	if request.Amount < 5000 {
		return nil, errs.NewVaildationError("amount at least 5,000")
	}
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewVaildationError("account type should be saving or checking")
	}

	acc := model.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now(),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Create(acc)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		ID:          newAcc.ID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}
	return &response, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	responses := []AccountResponse{}
	for _, account := range accounts {
		responses = append(responses, AccountResponse{
			Customer: CustomerResponse{
				ID:     account.Customer.ID,
				Name:   account.Customer.Name,
				Status: account.Customer.Status,
			},
			ID:          account.ID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}

	return responses, nil
}

func (s accountService) GetAccount(id int) ([]AccountResponse, error) {
	account, err := s.accRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	response := []AccountResponse{}
	response = append(response, AccountResponse{
		Customer: CustomerResponse{
			ID:     account.Customer.ID,
			Name:   account.Customer.Name,
			Status: account.Customer.Status,
		},
		ID:          account.ID,
		OpeningDate: account.OpeningDate,
		AccountType: account.AccountType,
		Amount:      account.Amount,
		Status:      account.Status,
	})

	return response, nil
}

func (s accountService) UpdateAccount(accountID int, request UpdateAccountRequest) (*AccountResponse, error) {

	// Validate
	if request.Amount < 5000 {
		return nil, errs.NewVaildationError("amount at least 5,000")
	}
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewVaildationError("account type should be saving or checking")
	}

	acc := model.Account{
		CustomerID:  request.CustomerID,
		OpeningDate: time.Now(),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Update(accountID, acc)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	response := AccountResponse{
		ID:          newAcc.ID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &response, nil
}

func (s accountService) DeleteAccount(accountID int) error {
	err := s.accRepo.Delete(accountID)
	if err != nil {
		logs.Error(err)
		return errs.NewUnexpectedError()
	}
	return nil
}
