package http

import (
	"net/http"

	"github.com/antonpodkur/Emstaht/internal/auth"
	"github.com/gin-gonic/gin"
)

type authHandlers struct{}

func NewAuthHandlers() auth.Handlers {
	return &authHandlers{}
}

func (h *authHandlers) Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello from controller"})
}
