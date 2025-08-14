package configuration

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"relazy/core/project"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Add your config fields here, e.g.:
	BasePath string            `yaml:"basePath"`
	Timeout  int               `yaml:"timeout"`
	Projects []project.Project `yaml:"projects"`
}

func LoadConfig() Config {
	var cfg Config
	err := cfg.Load()
	if err != nil {
		// Set defaults if loading fails
		cfg.BasePath = "$HOME/ws"
		cfg.Timeout = 30
	}
	return cfg
}
func configPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(usr.HomeDir, ".config", "relazy", "config.yaml"), nil
}

func (c *Config) Load() error {
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, c)
}

func (c *Config) Save() error {
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (c *Config) CreateProject(name, description, icon string, repos []project.ProjectRepository) (*project.Project, error) {
	repoRefs := c.mapReposToRefs(repos)
	appProject := &project.Project{
		Name:        name,
		Description: description,
		Icon:        icon,
		Repos:       repoRefs,
	}
	c.Projects = append(c.Projects, *appProject)
	err := c.Save()
	if err != nil {
		return nil, err
	}
	return appProject, nil
}
func (c *Config) mapReposToRefs(repos []project.ProjectRepository) []project.RepoRef {
	var repoRefs []project.RepoRef
	for _, repo := range repos {
		repoRefs = append(repoRefs, project.RepoRef{
			Owner: repo.Owner,
			Name:  repo.Name,
		})
	}
	return repoRefs
}
func (c *Config) GetAllProjects() ([]project.ProjectRef, error) {
	var projectRefs []project.ProjectRef
	for _, appProject := range c.Projects {
		projectRefs = append(projectRefs, project.ProjectRef{
			Name: appProject.Name,
			Icon: appProject.Icon,
		})
	}
	return projectRefs, nil
}

func (c *Config) Project(name string) (*project.Project, error) {
	for _, p := range c.Projects {
		if p.Name == name {
			return &p, nil
		}
	}
	return nil, errors.New("repoList not found")
}

func (c *Config) NewProject(name string) (project.Project, error) {
	project := project.Project{
		Name:        name,
		Description: "",
		Icon:        "",
		Repos:       []project.RepoRef{},
	}
	c.Projects = append(c.Projects, project)
	return project, nil

}

func (c *Config) GetProjectByName(name string) (*project.Project, error) {
	for _, appProject := range c.Projects {
		if appProject.Name == name {
			return &appProject, nil
		}
	}
	return nil, errors.New("project not found")
}
