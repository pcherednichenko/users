package handler

import (
	"go.uber.org/zap"

	"github.com/pcherednichenko/users/internal/database"
)

type server struct {
	l  *zap.SugaredLogger
	db database.DB
}

func NewServer(l *zap.SugaredLogger, db database.DB) *server {
	return &server{
		l:  l,
		db: db,
	}
}
