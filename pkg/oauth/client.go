package oauth

import "encoding/json"

type Client struct {
	Username string
	Password string
}

func NewClient(username, password string) (*Client, error) {
	return &Client{
		Username: "",
		Password: "",
	}, nil
}

func NewClientWithJSONPayload(payload []byte) (*Client, error) {
	var client *Client

	err := json.Unmarshal(payload, &client)
	if err != nil {
		return nil, err
	}

	return client, nil
}
