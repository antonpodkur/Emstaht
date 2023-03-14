package server

import (
	"github.com/gin-gonic/gin"
	authHttp "github.com/antonpodkur/Emstaht/internal/auth/delivery/http"
)

func (s *Server) MapHandlers(e *gin.Engine) error {
	// Init handlers
	authHandlers := authHttp.NewAuthHandlers()

	v1 := e.Group("/api/v1")

	authHttp.MapAuthRoutes(v1, authHandlers)
	return nil
}
