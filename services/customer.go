package services

import (
	"database/sql"
	"go-restful-api-template/errs"
	"go-restful-api-template/logs"
	"go-restful-api-template/repositories"
)

type customerService struct {
	customerRepo repositories.CustomerRepository
}

func NewCustomerService(customerRepo repositories.CustomerRepository) customerService {
	return customerService{customerRepo: customerRepo}
}

func (s *customerService) GetCustomers() ([]repositories.Customer, error) {
	customers, err := s.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return customers, nil
}

func (s *customerService) GetCustomer(id int) (*repositories.Customer, error) {
	customer, err := s.customerRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return customer, nil
}
