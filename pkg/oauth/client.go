package oauth

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
