package handler

import (
	"encoding/json"
	"go-restful-api-template/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(customerSrv services.CustomerService) CustomerHandler {
	return CustomerHandler{customerService: customerSrv}
}

func (h CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.customerService.GetCustomers()
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h CustomerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["Id"])
	customer, err := h.customerService.GetCustomer(customerID)
	if err != nil {
		handlerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
