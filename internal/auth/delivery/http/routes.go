package http

import (
	"github.com/antonpodkur/Emstaht/internal/auth"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroup *gin.RouterGroup, h auth.Handlers) {
	authGroup.POST("/register", h.Register)
	authGroup.POST("/login", h.Login)
}
