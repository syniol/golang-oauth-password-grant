package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
)

type Database struct {
	Ctx context.Context
	*sql.DB
}

var instance *Database

func NewDatabase() (*Database, error) {
	if instance != nil {
		return instance, nil
	}

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s",
		os.Getenv("DATABASE_USR"),
		os.Getenv("DATABASE_PWD"),
		"host.docker.internal",
	)
	cnn, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	instance = &Database{
		Ctx: context.Background(),
		DB:  cnn,
	}

	return instance, nil
}
