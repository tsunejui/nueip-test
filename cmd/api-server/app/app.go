package app

import (
	"context"
	"fmt"
	"nueip/cmd/api-server/conf"
	"nueip/cmd/api-server/models"
	"nueip/cmd/api-server/server"
)

type App struct {
	srv *server.Server
}

func New() *App {
	return &App{}
}

func (app *App) Start(ctx context.Context) error {
	config := conf.GetConfig()

	// Database
	dao, err := models.NewDAO(models.DBConfig{
		DBHost:     config.DBHost,
		DBPort:     config.DBPort,
		DBName:     config.DBName,
		DBUsername: config.DBUsername,
		DBPassword: config.DBPassword,
	})
	if err != nil {
		return fmt.Errorf("failed to connect to the db: %v", err)
	}

	// API Server
	s := server.New(dao)
	if err := s.WithPort(
		fmt.Sprintf(":%s", config.ListenPort),
	).
		WithContext(ctx).Start(); err != nil {
		return fmt.Errorf("failed to start the server: %v", err)
	}
	app.srv = s

	return nil
}

func (app *App) Stop(ctx context.Context) error {
	if err := app.srv.Stop(); err != nil {
		return fmt.Errorf("failed to stop the server: %v", err)
	}
	return nil
}
