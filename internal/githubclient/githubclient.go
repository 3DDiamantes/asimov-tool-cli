package githubclient

import (
	"asimov-tool-cli/internal/env"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const GithubApiUrl = "https://api.github.com"

type Client struct {
	AuthToken   string
	URL         string
	restyClient *resty.Client
}

func New() (*Client, error) {
	token, err := env.GetToken()
	if err != nil {
		return nil, err
	}

	return &Client{
		AuthToken:   token,
		restyClient: resty.New(),
	}, nil
}

func (c *Client) doPost(url string, body interface{}) (*resty.Response, error) {
	resp, err := c.restyClient.R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetAuthScheme("token").
		SetAuthToken(c.AuthToken).
		SetBody(body).
		Post(fmt.Sprintf("%s%s", GithubApiUrl, url))
	return resp, err
}
