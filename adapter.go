package main

import (
	"go-restful-api-template/handler"
	"go-restful-api-template/repository"
	"go-restful-api-template/service"

	"gorm.io/gorm"
)

func CustomerAdapter(db *gorm.DB) handler.CustomerHandler {
	customerRepository := repository.NewCustomerRepositoryImpl(db)
	// customerRepository := repositories.NewCustomerRepositoryMock() // Use Mock
	customerService := service.NewCustomerService(customerRepository)
	return handler.NewCustomerHandler(customerService)
}

func AccountAdapter(db *gorm.DB) handler.AccountHandler {
	accountRepository := repository.NewAccountRepositoryImpl(db)
	accountService := service.NewAccountService(accountRepository)
	return handler.NewAccountHandler(accountService)
}
