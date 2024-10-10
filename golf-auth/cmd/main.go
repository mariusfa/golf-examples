package main

import (
	appconfig "auth/internal/config"
	"auth/internal/controllers"
	"auth/internal/repositories"
	"fmt"
	"net/http"

	"github.com/mariusfa/golf/config"
	"github.com/mariusfa/golf/health"
	"github.com/mariusfa/golf/logging/accesslog"
	"github.com/mariusfa/golf/logging/applog"
	"github.com/mariusfa/golf/logging/middlewarelog"
	"github.com/mariusfa/golf/logging/tracelog"
	"github.com/mariusfa/golf/middleware"
)

const (
	APP_NAME = "auth"
	ENV_FILE = ".env"
)

func setupRouter() *http.ServeMux {
	router := http.NewServeMux()
	health.RegisterRoute(router)

	return router
}

func main() {
	applog.SetAppName(APP_NAME)
	accesslog.SetAppName(APP_NAME)
	tracelog.SetAppName(APP_NAME)
	middlewarelog.SetAppName(APP_NAME)
	var appConfig appconfig.Config
	err := config.GetConfig(ENV_FILE, &appConfig)
	if err != nil {
		panic(err)
	}

	router := setupRouter()

	userRepo := repositories.NewUserRepository()
	helloController := controllers.NewHelloController(userRepo)

	adminUser := middleware.User{Id: "123", Name: "admin"}
	userRepo.AddUser(adminUser)

	authParams := middleware.NewAuthParams("secret", userRepo, middlewarelog.GetLogger())
	helloController.RegisterRoutes(router, authParams)

	addr := fmt.Sprintf(":%s", appConfig.Port)
	applog.Info(fmt.Sprintf("Starting app %s on %s", APP_NAME, addr))
	http.ListenAndServe(addr, router)
}
