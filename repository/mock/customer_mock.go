package repositories

// import (
// 	"errors"
// 	"go-restful-api-template/repository"
// )

// type CustomerRepositoryMock struct {
// 	customers []repository.Customer
// }

// func NewCustomerRepositoryMock() CustomerRepositoryMock {
// 	customers := []repository.Customer{
// 		{ID: 1001, Name: "Mock-1", DateOfBirth: "2021-01-01 14:55:12", City: "Khonkaen"},
// 		{ID: 1002, Name: "Mock-2", DateOfBirth: "2021-02-02 14:55:12", City: "Bangkok"},
// 		{ID: 1003, Name: "Mock-3", DateOfBirth: "2021-03-03 14:55:12", City: "Ayuttaya"},
// 		{ID: 1004, Name: "Mock-4", DateOfBirth: "2021-04-04 14:55:12", City: "Lei"},
// 	}

// 	return CustomerRepositoryMock{customers: customers}
// }

// func (r CustomerRepositoryMock) GetAll() ([]repository.Customer, error) {
// 	return r.customers, nil
// }

// func (r CustomerRepositoryMock) GetById(id int) (*repository.Customer, error) {
// 	for _, customer := range r.customers {
// 		if customer.CustomerID == id {
// 			return &customer, nil
// 		}
// 	}
// 	return nil, errors.New("customer not found")
// }
