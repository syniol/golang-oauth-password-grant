package database

import (
	"context"
	"fmt"
	"os"

	"database/sql"
)

type Database struct {
	Ctx context.Context
	*sql.DB
}

var instance *Database

func NewDatabase(ctx context.Context) (*Database, error) {
	if instance != nil {
		return instance, nil
	}

	// https://www.postgresql.org/docs/current/libpq-ssl.html
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=require",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		func() string {
			if len(os.Getenv("DEBUG")) > 0 {
				return "127.0.0.1"
			}

			return "database"
		}(),
	)
	cnn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	instance = &Database{
		Ctx: ctx,
		DB:  cnn,
	}

	return instance, nil
}
