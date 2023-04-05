package clients

import (
	"oauth-password/pkg/database"
	"oauth-password/pkg/oauth"
)

type Repository struct {
	client *database.Database
}

func NewRepository() (*Repository, error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}

	return &Repository{
		client: db,
	}, nil
}

func (r *Repository) InsertSingle(
	clientCredential oauth.ClientCredential,
) (*Entity, error) {
	stmt, err := r.client.PrepareContext(
		r.client.Ctx,
		"INSERT INTO client_credential (data) VALUES (?)",
	)
	if err != nil {
		return nil, err
	}

	result, err := stmt.ExecContext(r.client.Ctx, clientCredential)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &Entity{
		ID:   uint(id),
		Data: clientCredential,
	}, nil
}
