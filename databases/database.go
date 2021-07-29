package databases

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

// DB is a global var for connect DB
var DB *sqlx.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig use for building DB config
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "0.0.0.0",
		Port:     3306,
		User:     "root",
		DBName:   "bank_database",
		Password: os.Getenv("DB_PASSWORD"),
	}
	return &dbConfig
}

// DbURL use for create DB connection URL
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
