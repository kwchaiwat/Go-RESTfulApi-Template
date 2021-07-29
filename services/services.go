package services

import "go-restful-api-template/models"

type BankService interface {
	GetBanks() ([]models.Bank, error)
	GetBank(id int) (*models.Bank, error)
}
