package database

import (
	"context"
	"database/sql"
	"fmt"
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

	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s/oauth?sslmode=disable",
		//os.Getenv("DATABASE_USR"),
		"oauth_usr",
		//os.Getenv("DATABASE_PWD"),
		"DummyPassword1",
		//"127.0.0.1",
		"host.docker.internal",
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
