package http

import (
	"github.com/antonpodkur/Emstaht/internal/middleware"
	"github.com/antonpodkur/Emstaht/internal/pages"
	"github.com/gin-gonic/gin"
)

func MapPagesRoutes(pagesGroup *gin.RouterGroup, h pages.Handlers, mw *middleware.MiddlewareManager) {
	pagesGroup.GET("/:userId", h.GetByUserId)
	pagesGroup.GET("/exps", h.GetWorkExperiences)
	pagesGroup.GET("/skills", h.GetSkills)
	pagesGroup.Use(mw.JwtAuthMiddleware())
	pagesGroup.POST("/", h.Create)
	pagesGroup.PUT("/", h.Update)
	pagesGroup.POST("/exp", h.CreateWorkExperience)
	pagesGroup.POST("/exps", h.CreateWorkExperiences)
	pagesGroup.PUT("/exp", h.UpdateWorkExperience)
	pagesGroup.DELETE("/exp/:expId", h.DeleteWorkExperience)
	pagesGroup.POST("/skill", h.CreateSkill)
	pagesGroup.POST("/skills", h.CreateSkills)
	pagesGroup.PUT("/skill", h.UpdateSkill)
	pagesGroup.DELETE("/skill/:skillId", h.DeleteSkill)
}
