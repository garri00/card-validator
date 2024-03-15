package main

import (
	"fmt"
	"net/http"
	"time"

	"card-validator/api"
	"card-validator/pkg/logger"
	"card-validator/src/config"
)

func main() {
	configs, err := config.GetConfig()
	if err != nil {
		logger.Log.Fatal().Err(err).Msg("failed to load .env")
	}

	logger.SetLogLevel(configs)

	logger.Log.Info().Msg("server started")

	routes := api.NewRouter()

	server := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", configs.Host, configs.Port),
		Handler:           routes,
		ReadHeaderTimeout: 3 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Log.Fatal().Msg("server crashed")
	}
}
