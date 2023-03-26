package middleware

import (
	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/auth"
)

type MiddlewareManager struct {
	authUsecase auth.Usecase
	cfg *config.Config
}

func NewMiddlewareManager(authUsecase auth.Usecase, cfg *config.Config) *MiddlewareManager {
	return &MiddlewareManager{
		authUsecase: authUsecase,
		cfg: cfg,
	}
}
