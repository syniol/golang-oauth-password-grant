package oauth

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Username string

func (u Username) Validate() error {
	if len(u) > 32 {
		return fmt.Errorf("maximum length of username is 32 characters")
	}

	if len(u) < 2 {
		return fmt.Errorf("minimum length of username is 2 characters")
	}

	return nil
}

func (u Username) String() string {
	return string(u)
}

type Password string

func (p Password) Validate() error {
	if len(p) > 128 {
		return fmt.Errorf("maximum length of password is 128 characters")
	}

	if len(p) < 8 {
		return fmt.Errorf("minimum length of password is 8 characters")
	}

	return nil
}

func (p Password) String() string {
	return string(p)
}

const (
	PasswordGrantFieldGrantType = "grant_type"
	PasswordGrantFieldUsername  = "username"
	PasswordGrantFieldPassword  = "password"
)

type PasswordGrantRequest struct {
	GrantType GrantType `json:"grant_type"`
	Username  Username  `json:"username"`
	Password  Password  `json:"password"`
}

func (pgr *PasswordGrantRequest) ToString() string {
	jsonFormat, _ := json.Marshal(pgr)

	return string(jsonFormat)
}

type PasswordGrantResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   uint   `json:"expires_in"`
}

func (pgr *PasswordGrantResponse) Byte(prettyJSON bool) []byte {
	if prettyJSON {
		output, _ := json.MarshalIndent(pgr, "", "\t")

		return output
	}

	output, _ := json.Marshal(pgr)

	return output
}

func (pgr *PasswordGrantResponse) String() string {
	jsonFormat, _ := json.Marshal(pgr)

	return string(jsonFormat)
}

func NewPasswordGrantResponse(token string) *PasswordGrantResponse {
	return &PasswordGrantResponse{
		AccessToken: token,
		TokenType:   TokenTypeBearer,
		ExpiresIn:   3600,
	}
}

// NewPasswordGrantRequest 127.0.0.1/oauth2/token
// POST /token HTTP/1.1
func NewPasswordGrantRequest(payload []byte) (*PasswordGrantRequest, error) {
	var req *PasswordGrantRequest

	err := json.Unmarshal(payload, &req)
	if err != nil {
		return nil, err
	}

	if req.GrantType != GrantTypePassword {
		return nil, fmt.Errorf(
			"grant type should be %q but %q is given",
			GrantTypePassword,
			req.GrantType,
		)
	}

	err = req.Username.Validate()
	if err != nil {
		return nil, err
	}

	err = req.Password.Validate()
	if err != nil {
		return nil, err
	}

	return req, nil
}

func NewPasswordGrantRequestWithForm(form url.Values) (*PasswordGrantRequest, error) {
	payloadMap := map[string]interface{}{
		PasswordGrantFieldGrantType: form.Get(PasswordGrantFieldGrantType),
		PasswordGrantFieldUsername:  form.Get(PasswordGrantFieldUsername),
		PasswordGrantFieldPassword:  form.Get(PasswordGrantFieldPassword),
	}

	payload, err := json.Marshal(payloadMap)
	if err != nil {
		return nil, err
	}

	return NewPasswordGrantRequest(payload)
}
