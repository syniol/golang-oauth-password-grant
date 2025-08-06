package clients

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

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
	ctx context.Context,
	clientCredential oauth.ClientCredential,
) (*Entity, error) {
	data, err := json.Marshal(clientCredential)
	if err != nil {
		return nil, err
	}

	_, err = r.client.ExecContext(
		ctx,
		`INSERT INTO public.client_credential (id, data) VALUES (DEFAULT, $1)`,
		data,
	)
	if err != nil {
		return nil, err
	}

	return &Entity{
		Data: clientCredential,
	}, nil
}

func (r *Repository) FindByUsername(ctx context.Context, username oauth.Username) (*Entity, error) {
	stmt, err := r.client.PrepareContext(
		ctx,
		`SELECT id, data FROM client_credential WHERE data->>'username' = $1;`,
	)
	if err != nil {
		return nil, err
	}

	var id uint
	var dataColumn []byte
	err = stmt.QueryRowContext(ctx, username.String()).Scan(&id, &dataColumn)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("unable to find username: %s", username)
	}

	var clientCredential oauth.ClientCredential
	err = json.Unmarshal(dataColumn, &clientCredential)
	if err != nil {
		return nil, err
	}

	return &Entity{
		ID:   id,
		Data: clientCredential,
	}, nil
}
