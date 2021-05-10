package main

import (
	"github.com/pcherednichenko/users/internal"
	"github.com/pcherednichenko/users/internal/config"
	"go.uber.org/zap"
)

//
// @title Users API
// @version 1.0
// @description This is example of creating a rest api for storage and user handling system with db
// @license.name Apache 2.0
// @host localhost:8080
// @BasePath /
func main() {
	// I personally really like zap logger, so I will use it here
	productionLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer func() { _ = productionLogger.Sync() }() // TODO: better to handle even such errors
	sugaredLogger := productionLogger.Sugar()

	cfg := config.LoadConfigFromEnv(sugaredLogger)
	err = internal.RunUsersService(cfg, sugaredLogger)
	if err != nil {
		panic(err)
	}
}