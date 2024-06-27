package main

import (
	"fmt"
	"net/http"
	appconfig "ping/internal/config"
	"ping/internal/controllers"

	"github.com/mariusfa/golf/config"
	"github.com/mariusfa/golf/health"
	"github.com/mariusfa/golf/logging/accesslog"
	"github.com/mariusfa/golf/logging/applog"
	"github.com/mariusfa/golf/metrics"
)

func appSetup() *http.ServeMux {
	ping := controllers.NewPing()
	router := http.NewServeMux()
	ping.RegisterRoutes(router)
	health.RegisterRoute(router)
	metrics.RegisterRoute(router)
	return router
}

const (
	APP_NAME = "ping"
	ENV_FILE = ".env"
)

func main() {
	applog.SetAppName(APP_NAME)
	accesslog.SetAppName(APP_NAME)

	router := appSetup()

	var appConfig appconfig.Config
	err := config.GetConfig(ENV_FILE, &appConfig)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%s", appConfig.Port)
	applog.Info(fmt.Sprintf("Starting app %s on %s", APP_NAME, addr))
	http.ListenAndServe(addr, router)
}
