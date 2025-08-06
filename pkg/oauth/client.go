package oauth

import (
	"encoding/json"
	"github.com/google/uuid"
)

type ClientRequest struct {
	Username Username `json:"username"`
	Password Password `json:"password"`
}

type ClientResponse struct {
	ClientID string `json:"client_id"`
}

func (cr *ClientResponse) Bytes(prettyJSON bool) []byte {
	if prettyJSON {
		output, _ := json.MarshalIndent(cr, "", "\t")

		return output
	}

	output, _ := json.Marshal(cr)

	return output
}

func (c *ClientRequest) String() string {
	res, _ := json.Marshal(c)

	return string(res)
}

// NewClientRequest Deprecated
func NewClientRequest(payload []byte) (*ClientRequest, error) {
	var client *ClientRequest

	err := json.Unmarshal(payload, &client)
	if err != nil {
		return nil, err
	}

	err = client.Username.Validate()
	if err != nil {
		return nil, err
	}

	err = client.Password.Validate()
	if err != nil {
		return nil, err
	}

	return client, nil
}

type ClientCredential struct {
	Credential

	ClientID string `json:"clientId"`
	Username string `json:"username"`
}

func NewClientCredential(
	credential *Credential,
	username string,
) (*ClientCredential, error) {
	return &ClientCredential{
		Credential: *credential,
		ClientID:   uuid.NewString(),
		Username:   username,
	}, nil
}
