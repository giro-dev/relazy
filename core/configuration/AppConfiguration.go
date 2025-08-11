package configuration

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Add your config fields here, e.g.:
	BasePath string    `yaml:"basePath"`
	Timeout  int       `yaml:"timeout"`
	Projects []Project `yaml:"projects"`
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
	return filepath.Join(usr.HomeDir, ".config", "relazy.yaml"), nil
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

func (c *Config) CreateProject(name, description, icon string, repos []RepoRef) (*Project, error) {
	project := &Project{
		Name:        name,
		Description: description,
		Icon:        icon,
		Repos:       repos,
	}
	// Here you can add logic to persist the project if needed
	return project, nil
}

func (c *Config) GetAllProjects() ([]ProjectRef, error) {
	var projectRefs []ProjectRef
	for _, project := range c.Projects {
		projectRefs = append(projectRefs, ProjectRef{
			Name: project.Name,
			Icon: project.Icon,
		})
	}
	return projectRefs, nil
}

func (c *Config) Project(name string) (*Project, error) {
	for _, p := range c.Projects {
		if p.Name == name {
			return &p, nil
		}
	}
	return nil, errors.New("Project not found")
}
