package app

import (
	"github.com/wavelaunch/pikachu-sql/api/internal/handlers/healthcheckhandlers"
	"github.com/wavelaunch/pikachu-sql/api/internal/handlers/taskhandlers"
	"github.com/wavelaunch/pikachu-sql/api/internal/logger"
	"github.com/wavelaunch/pikachu-sql/api/internal/ports"
	"github.com/wavelaunch/pikachu-sql/api/internal/repositories/dbsettingsrepository"
	"github.com/wavelaunch/pikachu-sql/api/internal/services/taskservice"
)

type repositories struct {
	DbSettingsRepository ports.DBSettingsRepository
}

type services struct {
	TaskService ports.TaskService
}

type handlers struct {
	HealthCheckHandlers ports.HealthCheckHandlers
	TaskHandlers        ports.TaskHandlers
}

type Application struct {
	Log          *logger.Logger
	Config       *ApplicationConfig
	Repositories *repositories
	Services     *services
	Handlers     *handlers
}

func New() *Application {
	app := &Application{
		Log:    logger.New(),
		Config: parseConfig(),
	}
	app.initRepositories()
	app.initServices()
	app.initHandlers()

	return app
}

func (a *Application) initRepositories() {
	a.Repositories = &repositories{
		DbSettingsRepository: dbsettingsrepository.NewJsonConfigFileLoader("./config/db_connections.json"),
	}
}

func (a *Application) initServices() {
	a.Services = &services{
		TaskService: taskservice.New(a.Repositories.DbSettingsRepository),
	}
}

func (a *Application) initHandlers() {
	a.Handlers = &handlers{
		HealthCheckHandlers: healthcheckhandlers.New(),
		TaskHandlers:        taskhandlers.New(a.Services.TaskService),
	}
}
