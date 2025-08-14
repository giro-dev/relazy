package repoRepository

import "relazy/core/project"

type RepoRepository struct {
	BasePath     string
	remoteClient RemoteClient
}

func (r RepoRepository) GetAllReposByProject(appProject project.Project) ([]string, error) {

	for _, repo := range appProject.Repos {
		repoRefs = append(repoRefs, project.RepoRef{
			Owner: repo.Owner,
			Name:  repo.Name,
		})
	}
}

// NewRepoRepository creates a new RepoRepository with the given base path
func NewRepoRepository(basePath string, factoryType RemoteClientFactoryType) *RepoRepository {
	return &RepoRepository{
		BasePath:     basePath,
		remoteClient: RemoteClientFactory(factoryType),
	}
}
