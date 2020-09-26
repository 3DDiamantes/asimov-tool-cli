package repository

import (
	"asimov-tool-cli/internal/env"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const (
	githubApiUrl = "https://api.github.com"
)

type createPRBody struct {
	Title string `json:"title"`
	Head  string `json:"head"`
	Base  string `json:"base"`
}

type GithubRepository interface {
	CreatePR(title string, headBranch string, baseBranch string) (*resty.Response, error)
}

type githubRepository struct {
	Repository string
	Owner      string
	authToken  string
	client     *resty.Client
}

func NewGithubRepository(owner string, repository string) GithubRepository {
	token, err := env.GetToken()
	if err != nil {
		panic("Error getting GitHub token")
	}

	return &githubRepository{
		Repository: repository,
		Owner:      owner,
		authToken:  token,
		client:     resty.New(),
	}
}

func (r *githubRepository) CreatePR(title string, headBranch string, baseBranch string) (*resty.Response, error) {
	url := fmt.Sprintf("/repos/%s/%s/pulls", r.Owner, r.Repository)

	body := createPRBody{
		Title: title,
		Head:  headBranch,
		Base:  baseBranch,
	}

	resp, err := r.client.R().
		SetHeader("Accept", "application/vnd.github.v3+json").
		SetAuthScheme("token").
		SetAuthToken(r.authToken).
		SetBody(body).
		Post(fmt.Sprintf("%s%s", githubApiUrl, url))

	return resp, err
}
