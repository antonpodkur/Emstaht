package http

import (
	"github.com/antonpodkur/Emstaht/internal/middleware"
	"github.com/antonpodkur/Emstaht/internal/pages"
	"github.com/gin-gonic/gin"
)

func MapPagesRoutes(pagesGroup *gin.RouterGroup, h pages.Handlers, mw *middleware.MiddlewareManager) {
	pagesGroup.Use(mw.JwtAuthMiddleware())
	pagesGroup.GET("/:userId", h.GetByUserId)
	pagesGroup.POST("/", h.Create)
	pagesGroup.PUT("/", h.Update)
}
