package localRepository

import (
	"os"
	"path/filepath"
)

type LocalRepository struct {
	Basepath string
}

func (repository *LocalRepository) GetAllReposByOwner(owner string) ([]string, error) {
	var repos []string
	baseDir := filepath.Join(repository.Basepath, owner)
	err := filepath.Walk(baseDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && filepath.Base(path) == ".git" {
			repos = append(repos, filepath.Dir(path))
		}
		return nil
	})
	return repos, err

}
