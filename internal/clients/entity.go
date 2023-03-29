package clients

import "oauth-password/pkg/oauth"

type Entity struct {
	ID   uint                   `json:"id"`
	Data oauth.ClientCredential `json:"data"`
}
