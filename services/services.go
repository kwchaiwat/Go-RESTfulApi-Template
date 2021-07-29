package services

import "go-restful-api-template/repositories"

type CustomerService interface {
	GetCustomers() ([]repositories.Customer, error)
	GetCustomer(id int) (*repositories.Customer, error)
}
