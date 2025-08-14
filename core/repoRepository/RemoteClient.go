package repoRepository

import "github.com/google/go-github/v74/github"

type RemoteClient interface {
	getRepo(owner string, repoName string) (*github.Repository, error)
}
type RemoteClientFactoryType int

const (
	Github RemoteClientFactoryType = iota
)

var RemoteClientFactory = func(factoryType RemoteClientFactoryType) RemoteClient {
	switch factoryType {
	case Github:
		return NewGithubRemoteClient()
	default:
		return nil
	}
}
