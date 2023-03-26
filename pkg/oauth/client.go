package oauth

import "encoding/json"

type Client struct {
	Username Username `json:"username"`
	Password Password `json:"password"`
}

func (c *Client) String() string {
	res, _ := json.Marshal(c)

	return string(res)
}

func NewClient(payload []byte) (*Client, error) {
	var client *Client

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
