package main

import (
	"fmt"
	"go-restful-api-template/databases"
	"go-restful-api-template/handler"
	"go-restful-api-template/repositories"
	"go-restful-api-template/services"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {

	// InitConfig
	initTimeZone()
	db := initDatabase()

	// Plug Adapter
	bankRepository := repositories.NewBankRepositoryImpl(db)
	// bankRepository := repositories.NewBankRepositoryMock() // Use Mock
	bankService := services.NewBankService(bankRepository)
	bankHandler := handler.NewBankHandler(&bankService)

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/banks", bankHandler.GetBanks).Methods(http.MethodGet)
	router.HandleFunc("/banks/{bankId:[0-9]+}", bankHandler.GetBank).Methods(http.MethodGet)

	// ListenAndServe PORT 8000
	http.ListenAndServe(":8000", router)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
	db, err := sqlx.Open("mysql", databases.DbURL(databases.BuildDBConfig()))
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
