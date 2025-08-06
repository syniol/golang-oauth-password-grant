package database

import (
	"fmt"
	"os"
	"sync"

	"database/sql"
)

type Database struct {
	*sql.DB
}

var once sync.Once

func DBConnection() (postgresInstance *sql.DB, dbError error) {
	postgresInstance, dbError = sql.Open("postgres", fmt.Sprintf(
		"host=%s port=5432 user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	))
	if dbError != nil {
		return
	}

	postgresInstance.SetMaxOpenConns(1)

	return postgresInstance, dbError
}

var cnn *sql.DB

func NewDatabase() (*Database, error) {
	var dbError error
	once.Do(func() {
		cnn, dbError = DBConnection()
	})

	return &Database{
		DB: cnn,
	}, dbError
}
