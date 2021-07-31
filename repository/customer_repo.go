package repository

import model "go-restful-api-template/models"

type CustomerRepository interface {
	GetAll() ([]model.Customer, error)
	GetById(int) (*model.Customer, error)
}
