package main

import (
	appconfig "auth/internal/config"
	"fmt"
	"net/http"

	"github.com/mariusfa/golf/config"
	"github.com/mariusfa/golf/health"
	"github.com/mariusfa/golf/logging/applog"
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
	var appConfig appconfig.Config
	err := config.GetConfig(ENV_FILE, &appConfig)
	if err != nil {
		panic(err)
	}

	router := setupRouter()

	addr := fmt.Sprintf(":%s", appConfig.Port)
	applog.Info(fmt.Sprintf("Starting app %s on %s", APP_NAME, addr))
	http.ListenAndServe(addr, router)
}
