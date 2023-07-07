package handlers

import (
	"github.com/DeneesK/go-mart/internal/api-go-mart/repository"
	"go.uber.org/zap"
)

type Handler struct {
	db  repository.Storage
	log *zap.SugaredLogger
}

func New(db repository.Storage, log *zap.SugaredLogger) *Handler {
	return &Handler{
		db:  db,
		log: log,
	}
}
