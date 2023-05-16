package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"oauth-password/pkg/database"
	"oauth-password/pkg/oauth"
)

type Repository struct {
	client *database.Database
}

func NewRepository(ctx context.Context) (*Repository, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	db, err := database.NewDatabase(ctx)
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
	data, err := json.Marshal(clientCredential)
	if err != nil {
		return nil, err
	}

	_, err = r.client.ExecContext(
		r.client.Ctx,
		fmt.Sprintf(
			`INSERT INTO public.client_credential (id, data) VALUES (DEFAULT, '%s')`,
			data,
		),
	)
	if err != nil {
		return nil, err
	}

	return &Entity{
		Data: clientCredential,
	}, nil
}

func (r *Repository) FindByUsername(username oauth.Username) (*Entity, error) {
	rows, err := r.client.QueryContext(
		r.client.Ctx,
		fmt.Sprintf(
			`SELECT id, data FROM public.client_credential WHERE data->>'username' = '%s';`,
			username.String(),
		),
	)
	if err != nil {
		return nil, err
	}

	var id uint
	var dataColumn []byte
	for rows.Next() {
		err = rows.Scan(&id, &dataColumn)
		if err != nil {
			return nil, err
		}
	}

	if len(dataColumn) == 0 {
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
