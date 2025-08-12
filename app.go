package main

import (
	"context"
	"relazy/core/configuration"
	"relazy/core/localRepository"
)

// App struct
type App struct {
	ctx       context.Context
	config    configuration.Config
	localRepo localRepository.LocalRepository
}

// NewApp creates a new App application struct
func NewApp() *App {
	config := configuration.LoadConfig()
	_ = config.Save()
	return &App{
		config:    config,
		localRepo: localRepository.LocalRepository{Basepath: config.BasePath},
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) ListLocalRepos(owner string) ([]string, error) {
	return a.localRepo.GetAllReposByOwner(owner)
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
