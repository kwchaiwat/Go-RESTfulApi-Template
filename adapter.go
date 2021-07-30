package main

import (
	"go-restful-api-template/handler"
	"go-restful-api-template/repository"
	"go-restful-api-template/service"

	"github.com/jmoiron/sqlx"
)

func CustomerAdapter(db *sqlx.DB) handler.CustomerHandler {
	customerRepository := repository.NewCustomerRepositoryImpl(db)
	// customerRepository := repositories.NewCustomerRepositoryMock() // Use Mock
	customerService := service.NewCustomerService(customerRepository)
	return handler.NewCustomerHandler(customerService)
}

func AccountAdapter(db *sqlx.DB) handler.AccountHandler {
	accountRepository := repository.NewAccountRepositoryImpl(db)
	accountService := service.NewAccountService(accountRepository)
	return handler.NewAccountHandler(accountService)
}
