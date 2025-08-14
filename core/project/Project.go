package project

type RepoRef struct {
	Owner string `yaml:"owner"`
	Name  string `yaml:"name"`
}

type ProjectRef struct {
	Name string `yaml:"name"`
	Icon string `yaml:"icon"`
}

type Project struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Icon        string    `yaml:"icon"`
	Repos       []RepoRef `yaml:"repos"`
}
