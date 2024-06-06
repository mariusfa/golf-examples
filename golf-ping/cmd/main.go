package main

import (
	"fmt"
	"net/http"
	appconfig "ping/internal/config"
	"ping/internal/controllers"

	"github.com/mariusfa/golf/config"
	"github.com/mariusfa/golf/health"
	accesslog "github.com/mariusfa/golf/logging/access-log"
	applog "github.com/mariusfa/golf/logging/app-log"
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
	applog.AppLog = applog.NewAppLogger(APP_NAME)
	accesslog.AccessLog = accesslog.NewAccessLogger(APP_NAME)
	
	router := appSetup()

	var appConfig appconfig.Config
	err := config.GetConfig(ENV_FILE, &appConfig)
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%s", appConfig.Port)
	applog.AppLog.Info(fmt.Sprintf("Starting app %s on %s", APP_NAME, addr))
	http.ListenAndServe(addr, router)
}
