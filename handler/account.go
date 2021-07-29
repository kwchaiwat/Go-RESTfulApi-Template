package handler

import (
	"encoding/json"
	"go-restful-api-template/errs"
	"go-restful-api-template/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {

	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])

	if r.Header.Get("content-type") != "application/json" {
		handlerError(w, errs.NewVaildationError("request body incorrect format"))
		return
	}

	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handlerError(w, errs.NewVaildationError("request body incorrect format"))
		return
	}

	res, err := h.accSrv.NewAccount(customerID, request)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["customerID"])
	res, err := h.accSrv.GetAccounts(customerID)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}
