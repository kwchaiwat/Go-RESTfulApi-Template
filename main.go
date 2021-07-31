package main

import (
	"fmt"
	"go-restful-api-template/logs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	initTimeZone()
	db := initDatabase()
	// db.AutoMigrate(repository.Customer{}, repository.Account{})

	// Plug Adapter
	customerHandler := CustomerAdapter(db)
	accountHandler := AccountAdapter(db)

	app := fiber.New()
	app.Use(requestid.New(), logger.New(logger.Config{TimeZone: "Asia/Bangkok"}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	v1 := app.Group("/api/v1", func(c *fiber.Ctx) error {
		c.Set("Version", "v1")
		return c.Next()
	})

	//Router
	v1.Get("/customers", customerHandler.GetCustomers)
	v1.Get("/customers/:customerID", customerHandler.GetCustomer)
	v1.Get("/customers/:customerID/:accounts?", accountHandler.GetAccounts)
	v1.Post("/customers/:customerID/:accounts?", accountHandler.NewAccount)

	// ListenAndServe PORT 8000
	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	app.Listen(fmt.Sprintf("%v:%v", viper.GetString("app.host"), viper.GetInt("app.port")))
}
