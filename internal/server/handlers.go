package server

import (
	authHttp "github.com/antonpodkur/Emstaht/internal/auth/delivery/http"
	authRepo "github.com/antonpodkur/Emstaht/internal/auth/repository"
	authUseCase "github.com/antonpodkur/Emstaht/internal/auth/usecase"
	"github.com/antonpodkur/Emstaht/internal/middleware"
	pagesHttp "github.com/antonpodkur/Emstaht/internal/pages/delivery/http"
	pagesRepo "github.com/antonpodkur/Emstaht/internal/pages/repository"
	pagesUseCase "github.com/antonpodkur/Emstaht/internal/pages/usecase"
	"github.com/gin-gonic/gin"
)

func (s *Server) MapHandlers(e *gin.Engine) error {
	//Init repositories
	authRepository := authRepo.NewAuthRepository(s.db)
	pagesRepository := pagesRepo.NewPagesRepository(s.db)

	// Init UseCases
	authUsecase := authUseCase.NewAuthUsecase(s.cfg, authRepository)
	pagesUsecase := pagesUseCase.NewPagesUsecase(pagesRepository)

	// Init handlers
	authHandlers := authHttp.NewAuthHandlers(s.cfg, authUsecase)
	pagesHandlers := pagesHttp.NewPagesHandlers(pagesUsecase)

	mw := middleware.NewMiddlewareManager(authUsecase, s.cfg)

	v1 := e.Group("/api/v1")

	authGroup := v1.Group("/auth")
	pagesGroup := v1.Group("/pages")

	authHttp.MapAuthRoutes(authGroup, authHandlers, mw)
	pagesHttp.MapPagesRoutes(pagesGroup, pagesHandlers, mw)
	return nil
}
