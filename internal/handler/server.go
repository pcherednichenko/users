package handler

import (
	"github.com/pcherednichenko/users/internal/database"
	"go.uber.org/zap"
)

type server struct {
	l  *zap.SugaredLogger
	db database.DB
}

func NewServer(l *zap.SugaredLogger, db database.DB) *server {
	return &server{
		l: l,
		db: db,
	}
}
