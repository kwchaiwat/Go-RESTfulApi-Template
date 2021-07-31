package handler

import (
	"fmt"
	"go-restful-api-template/service"

	"github.com/gofiber/fiber/v2"
)

type AccountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) AccountHandler {
	return AccountHandler{accSrv: accSrv}
}

func (h AccountHandler) NewAccount(c *fiber.Ctx) error {
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

func (h AccountHandler) GetAccounts(c *fiber.Ctx) error {
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

func (h AccountHandler) GetAccount(c *fiber.Ctx) error {
	fmt.Printf("IsJson: %v\n", c.Is("json"))
	accountID, err := c.ParamsInt("accountID")
	if err != nil {
		return fiber.ErrBadRequest
	}
	res, err := h.accSrv.GetAccount(accountID)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (h AccountHandler) UpdateAccount(c *fiber.Ctx) error {
	accountID, err := c.ParamsInt("accountID")
	if err != nil {
		return fiber.ErrBadRequest
	}
	account := service.UpdateAccountRequest{}
	err = c.BodyParser(&account)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "request body incorrect format")
	}
	res, err := h.accSrv.UpdateAccount(accountID, account)
	if err != nil {
		return err
	}
	return c.JSON(res)
}

func (h AccountHandler) DeleteAccount(c *fiber.Ctx) error {
	accountID, err := c.ParamsInt("accountID")
	if err != nil {
		return fiber.ErrBadRequest
	}
	err = h.accSrv.DeleteAccount(accountID)
	if err != nil {
		return err
	}
	return nil
}
