package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

// TODO: fetch from env vars
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "productionService"
)

func SetupDB() *sql.DB {
	HOST := os.Getenv("HOST")
	if HOST == "" {
		HOST = "localhost"
	}
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s", DB_USER, DB_PASSWORD, DB_NAME, HOST)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return nil
	}

	return db
}
