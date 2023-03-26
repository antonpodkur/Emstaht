package http

import (
	"github.com/antonpodkur/Emstaht/internal/auth"
	"github.com/antonpodkur/Emstaht/internal/middleware"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handlers, mw *middleware.MiddlewareManager) {
	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
	authGroup.Use(mw.JwtAuthMiddleware())
	authGroup.GET("/me", h.Me)
}
