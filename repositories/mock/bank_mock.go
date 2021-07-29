package repositories

import (
	"errors"
	"go-restful-api-template/repositories"
)

type BankRepositoryMock struct {
	banks []repositories.Bank
}

func NewBankRepositoryMock() BankRepositoryMock {
	banks := []repositories.Bank{
		{Id: 1001, AccountNumber: "Mock-1", Trust: 111.99, TransactionFee: 10},
		{Id: 1002, AccountNumber: "Mock-2", Trust: 222.99, TransactionFee: 20},
		{Id: 1003, AccountNumber: "Mock-3", Trust: 333.99, TransactionFee: 30},
		{Id: 1004, AccountNumber: "Mock-4", Trust: 444.99, TransactionFee: 40},
	}

	return BankRepositoryMock{banks: banks}
}

func (r BankRepositoryMock) GetAll() ([]repositories.Bank, error) {
	return r.banks, nil
}

func (r BankRepositoryMock) GetById(id int) (*repositories.Bank, error) {
	for _, bank := range r.banks {
		if bank.Id == id {
			return &bank, nil
		}
	}
	return nil, errors.New("bank not found")
}
