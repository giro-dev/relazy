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
