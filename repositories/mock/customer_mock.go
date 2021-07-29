package repositories

import (
	"errors"
	"go-restful-api-template/repositories"
)

type CustomerRepositoryMock struct {
	customers []repositories.Customer
}

func NewCustomerRepositoryMock() CustomerRepositoryMock {
	customers := []repositories.Customer{
		{Id: 1001, AccountNumber: "Mock-1", Trust: 111.99, TransactionFee: 10},
		{Id: 1002, AccountNumber: "Mock-2", Trust: 222.99, TransactionFee: 20},
		{Id: 1003, AccountNumber: "Mock-3", Trust: 333.99, TransactionFee: 30},
		{Id: 1004, AccountNumber: "Mock-4", Trust: 444.99, TransactionFee: 40},
	}

	return CustomerRepositoryMock{customers: customers}
}

func (r CustomerRepositoryMock) GetAll() ([]repositories.Customer, error) {
	return r.customers, nil
}

func (r CustomerRepositoryMock) GetById(id int) (*repositories.Customer, error) {
	for _, customer := range r.customers {
		if customer.Id == id {
			return &customer, nil
		}
	}
	return nil, errors.New("customer not found")
}
