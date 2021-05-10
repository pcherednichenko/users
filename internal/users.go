package internal

import (
	"net/http"

	"github.com/pcherednichenko/users/internal/config"
	"github.com/pcherednichenko/users/internal/database"
	"github.com/pcherednichenko/users/internal/handler"
	"go.uber.org/zap"
)

// RunUsersService starts the main service with http handler and connection
func RunUsersService(cfg config.Config, l *zap.SugaredLogger) error {
	postgres, err := database.NewPostgresConnection(cfg)
	if err != nil {
		return err
	}

	s := handler.NewServer(l, postgres)
	router := s.Router()

	// todo handle errors
	l.Infof("Starting service on port: %s", cfg.Port)
	err = http.ListenAndServe(cfg.Port, router)
	if err != nil {
		return err
	}
	return nil
}