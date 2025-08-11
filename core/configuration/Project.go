package configuration

type RepoRef struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

type ProjectRef struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}

type Project struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Repos       []RepoRef `json:"repos"`
}
