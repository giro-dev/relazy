package main

import (
	"context"
	"relazy/core/configuration"
	"relazy/core/repoRepository"
)

// App struct
type App struct {
	ctx       context.Context
	config    configuration.Config
	localRepo repoRepository.RepoRepository
}

// NewApp creates a new App application struct
func NewApp() *App {
	config := configuration.LoadConfig()
	_ = config.Save()
	return &App{
		config:    config,
		localRepo: repoRepository.RepoRepository{BasePath: config.BasePath},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ListProjectRepos(projectName string) ([]string, error) {
	project, err := a.config.GetProjectByName(projectName)
	if err != nil {
		return nil, err
	}
	if project == nil {
		return nil, configuration.ErrProjectNotFound
	}
	return a.localRepo.GetAllReposByProject(project)

}

func (a *App) GetProjects() ([]configuration.ProjectRef, error) {
	return a.config.GetAllProjects()
}

func (a *App) NewProject(name string) (*configuration.ProjectRef, error) {
	project, err := a.config.NewProject(name)
	if err != nil {
		return nil, err
	}
	if err := a.config.Save(); err != nil {
		return nil, err
	}
	return &configuration.ProjectRef{
		Name: project.Name,
		Icon: project.Icon,
	}, nil
}
