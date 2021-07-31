package repository

import model "go-restful-api-template/models"

type AccountRepository interface {
	Create(model.Account) (*model.Account, error)
	GetAll(int) ([]model.Account, error)
}
