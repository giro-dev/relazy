package repoRepository

import (
	"context"
	"os"

	"github.com/google/go-github/v74/github"
)

type GithubRemoteClient struct {
	client *github.Client
}

func (c *GithubRemoteClient) getRepo(owner string, repoName string) (*github.Repository, error) {
	repository, _, err := c.client.Repositories.Get(context.Background(), owner, repoName)
	if err != nil {
		return nil, err
	}
	return repository, nil
}

func NewGithubRemoteClient() RemoteClient {
	client, err := github.NewClient(nil).WithEnterpriseURLs("https://git.i.mercedes-benz.com", "https://git.i.mercedes-benz.com")
	if err != nil {
		panic("Failed to create GitHub client: " + err.Error())
	}

	remoteClient := GithubRemoteClient{
		client: client.WithAuthToken(os.Getenv("GITHUB_TOKEN")),
	}
	return &remoteClient
}
