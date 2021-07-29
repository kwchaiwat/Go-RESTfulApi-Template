package services

import (
	"database/sql"
	"go-restful-api-template/errs"
	"go-restful-api-template/logs"
	"go-restful-api-template/repositories"
)

type bankService struct {
	bankRepo repositories.BankRepository
}

func NewBankService(bankRepo repositories.BankRepository) bankService {
	return bankService{bankRepo: bankRepo}
}

func (s *bankService) GetBanks() ([]repositories.Bank, error) {
	banks, err := s.bankRepo.GetAll()
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return banks, nil
}

func (s *bankService) GetBank(id int) (*repositories.Bank, error) {
	bank, err := s.bankRepo.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("bank not found")
		}
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	return bank, nil
}
