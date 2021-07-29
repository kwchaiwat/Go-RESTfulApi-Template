package main

import (
	"fmt"
	"go-restful-api-template/databases"
	"go-restful-api-template/handler"
	"go-restful-api-template/repositories"
	"go-restful-api-template/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", databases.DbURL(databases.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	defer db.Close()
	// bankRepository := repositories.NewBankRepositoryImpl(db)
	bankRepository := repositories.NewBankRepositoryMock() // Use Mock
	bankService := services.NewBankService(bankRepository)
	bankHandler := handler.NewBankHandler(&bankService)

	router := mux.NewRouter()

	router.HandleFunc("/banks", bankHandler.GetBanks).Methods(http.MethodGet)
	router.HandleFunc("/banks/{bankId:[0-9]+}", bankHandler.GetBank).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)

}
