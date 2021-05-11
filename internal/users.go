package internal

import (
	"net/http"

	"go.uber.org/zap"

	"github.com/pcherednichenko/users/internal/config"
	"github.com/pcherednichenko/users/internal/database"
	"github.com/pcherednichenko/users/internal/handler"
)

// RunUsersService starts the main service with http handler and connection
func RunUsersService(cfg config.Config, l *zap.SugaredLogger) error {
	postgres, err := database.NewPostgresConnection(cfg)
	if err != nil {
		return err
	}

	s := handler.NewServer(l, postgres)
	router := s.Router()

	l.Infof("Starting service on port: %s", cfg.Port)
	err = http.ListenAndServe(cfg.Port, router)
	if err != nil {
		return err
	}
	return nil
}
