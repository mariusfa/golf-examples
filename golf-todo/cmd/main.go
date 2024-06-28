package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"todo/internal/controllers"
	"todo/internal/repositories"
	"todo/internal/services"

	appconfig "todo/internal/config"

	"github.com/mariusfa/golf/config"
	"github.com/mariusfa/golf/database"
	"github.com/mariusfa/golf/logging/accesslog"
	"github.com/mariusfa/golf/logging/applog"
)

func setupRouter(todoRepository services.TodoRepository) *http.ServeMux {
	router := http.NewServeMux()

	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)
	todoController.RegisterRoutes(router)

	return router
}

func setupDb(dbConfig *database.DbConfig) (*sql.DB, error) {
	db, err := database.Setup(dbConfig)
	if err != nil {
		return nil, err
	}

	err = database.Migrate(dbConfig, "migrations")
	if err != nil {
		return nil, err
	}

	return db, nil
}

const (
	APP_NAME = "todo"
	ENV_FILE = ".env"
)

func main() {
	applog.SetAppName(APP_NAME)
	accesslog.SetAppName(APP_NAME)

	var appConfig appconfig.Config
	err := config.GetConfig(ENV_FILE, &appConfig)
	if err != nil {
		panic(err)
	}

	dbConfig := appConfig.ToDbConfig()

	if dbConfig.StartupLocal == "true" {
		containerCleanUp, err := database.SetupContainer(dbConfig)
		if err != nil {
			panic(err)
		}
		defer containerCleanUp()
	}

	db, err := setupDb(dbConfig)
	if err != nil {
		panic(err)
	}

	todoRepository := repositories.NewTodoRepository(db)
	router := setupRouter(todoRepository)

	addr := fmt.Sprintf(":%s", appConfig.Port)
	applog.Info(fmt.Sprintf("Starting app %s on %s", APP_NAME, addr))
	http.ListenAndServe(addr, router)
}
