package repositories

import (
	"errors"
	"go-restful-api-template/models"
)

type BankRepositoryMock struct {
	banks []models.Bank
}

func NewBankRepositoryMock() BankRepositoryMock {
	banks := []models.Bank{
		{Id: 1001, AccountNumber: "Mock-1", Trust: 111.99, TransactionFee: 10},
		{Id: 1002, AccountNumber: "Mock-2", Trust: 222.99, TransactionFee: 20},
		{Id: 1003, AccountNumber: "Mock-3", Trust: 333.99, TransactionFee: 30},
		{Id: 1004, AccountNumber: "Mock-4", Trust: 444.99, TransactionFee: 40},
	}

	return BankRepositoryMock{banks: banks}
}

func (r BankRepositoryMock) GetAll() ([]models.Bank, error) {
	return r.banks, nil
}

func (r BankRepositoryMock) GetById(id int) (*models.Bank, error) {
	for _, bank := range r.banks {
		if bank.Id == id {
			return &bank, nil
		}
	}
	return nil, errors.New("bank not found")
}
