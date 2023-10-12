package repository

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func NewConnectMysqlDb() *sqlx.DB {
	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")

	parseTime := "true"
	loc := "Asia%2FJakarta"

	// conStr only for mysql connection
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		parseTime,
		loc,
	)

	db, err := sqlx.Open(dbConnection, conStr)
	if err != nil {
		fmt.Println("openDB", err.Error())
	}

	return db
}
