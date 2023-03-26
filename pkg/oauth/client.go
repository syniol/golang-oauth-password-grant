package oauth

import "encoding/json"

type Client struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	return client, nil
}
