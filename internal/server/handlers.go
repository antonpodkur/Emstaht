package server

import (
	authHttp "github.com/antonpodkur/Emstaht/internal/auth/delivery/http"
	authRepo "github.com/antonpodkur/Emstaht/internal/auth/repository"
	authUseCase "github.com/antonpodkur/Emstaht/internal/auth/usecase"
	"github.com/antonpodkur/Emstaht/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(e *gin.Engine) error {
	//Init repositories
	authRepository := authRepo.NewAuthRepository(s.db)

	// Init UseCases
	authUsecase := authUseCase.NewAuthUsecase(s.cfg, authRepository)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUsecase)

	mw := middleware.NewMiddlewareManager(authUsecase, s.cfg)

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)
	return nil
}
