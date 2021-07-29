package main

import (
	"fmt"
	"go-restful-api-template/handler"
	"go-restful-api-template/repositories"
	"go-restful-api-template/services"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
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
	log.Printf("Banking service started at port: %v", viper.GetInt("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		fmt.Println("statuse: ", err)
	}
	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
