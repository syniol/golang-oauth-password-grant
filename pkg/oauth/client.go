package oauth

import "encoding/json"

type ClientRequest struct {
	Username Username `json:"username"`
	Password Password `json:"password"`
}

func (c *ClientRequest) String() string {
	res, _ := json.Marshal(c)

	return string(res)
}

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
