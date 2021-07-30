package handler

import (
	"fmt"
	"go-restful-api-template/service"

	"github.com/gofiber/fiber/v2"
)

type CustomerHandler struct {
	customerService service.CustomerService
}

func NewCustomerHandler(customerSrv service.CustomerService) CustomerHandler {
	return CustomerHandler{customerService: customerSrv}
}

func (h CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	fmt.Printf("IsJson: %v\n", c.Is("json"))
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		return err
	}
	return c.JSON(customers)
}

func (h CustomerHandler) GetCustomer(c *fiber.Ctx) error {
	customerID, err := c.ParamsInt("customerID")
	if err != nil {
		return fiber.ErrBadRequest
	}

	customer, err := h.customerService.GetCustomer(customerID)
	if err != nil {
		return err
	}
	return c.JSON(customer)
}
