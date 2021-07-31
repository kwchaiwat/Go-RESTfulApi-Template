package repository

import model "go-restful-api-template/models"

type CustomerRepository interface {
	GetAll() ([]model.Customer, error)
	GetById(int) (*model.Customer, error)
}

type AccountRepository interface {
	Create(model.Account) (*model.Account, error)
	GetAll(int) ([]model.Account, error)
}
