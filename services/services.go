package services

import "go-restful-api-template/repositories"

type BankService interface {
	GetBanks() ([]repositories.Bank, error)
	GetBank(id int) (*repositories.Bank, error)
}
