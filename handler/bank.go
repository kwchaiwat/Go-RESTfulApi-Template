package handler

import (
	"encoding/json"
	"fmt"
	"go-restful-api-template/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type BankHandler struct {
	bankService services.BankService
}

func NewBankHandler(bankSrv services.BankService) BankHandler {
	return BankHandler{bankService: bankSrv}
}

func (h BankHandler) GetBanks(w http.ResponseWriter, r *http.Request) {
	banks, err := h.bankService.GetBanks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(banks)
}

func (h BankHandler) GetBank(w http.ResponseWriter, r *http.Request) {

	bankId, _ := strconv.Atoi(mux.Vars(r)["bankId"])
	bank, err := h.bankService.GetBank(bankId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bank)
}
