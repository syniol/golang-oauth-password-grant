package database

type Client interface{}

type Database struct {
	client Client
}

func (d *Database) Query(tableName string) {}

func NewDatabase() *Database {
	return &Database{client: nil}
}
