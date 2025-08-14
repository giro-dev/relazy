package project

type ProjectRepository struct {
	BasePath      string `yaml:"basePath"`
	Owner         string `yaml:"owner"`
	Name          string `yaml:"name"`
	RemoteURL     string `yaml:"remoteURL"`
	CurrentBranch string `yaml:"currentBranch"`
}
