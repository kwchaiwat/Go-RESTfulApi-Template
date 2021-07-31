package service

import (
	"database/sql"
	"go-restful-api-template/errs"
	"go-restful-api-template/logs"
	"go-restful-api-template/repository"
)

type CustomerResponse struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}

type customerService struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerService(customerRepo repository.CustomerRepository) customerService {
	return customerService{customerRepo: customerRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			ID:     customer.ID,
			Name:   customer.Name,
			Status: customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		ID:     customer.ID,
		Name:   customer.Name,
		Status: customer.Status,
	}

	return &custResponse, nil
}
