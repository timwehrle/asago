package api

import "net/http"

const (
	Base = "https://app.asana.com/api/1.0"
)

type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

func NewClient(token string) *Client {
	return &Client{
		BaseURL:    Base,
		HTTPClient: &http.Client{},
		Token:      token,
	}
}

func (c *Client) New(method, endpoint string) (*http.Request, error) {
	req, err := http.NewRequest(method, c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	return req, nil
}
