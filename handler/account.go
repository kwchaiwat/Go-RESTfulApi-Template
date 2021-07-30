package handler

import (
	"fmt"
	"go-restful-api-template/service"

	"github.com/gofiber/fiber/v2"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(c *fiber.Ctx) error {
	customerID, err := c.ParamsInt("customerID")
	if err != nil {
		return fiber.ErrBadRequest
	}
	account := service.NewAccountRequest{}
	err = c.BodyParser(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "request body incorrect format")
	}
	res, err := h.accSrv.NewAccount(customerID, account)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (h accountHandler) GetAccounts(c *fiber.Ctx) error {
	fmt.Printf("IsJson: %v\n", c.Is("json"))
	customerID, err := c.ParamsInt("customerID")
	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		return err
	}
	return c.JSON(res)
}
