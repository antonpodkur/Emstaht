package pages

import (
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetByUserId(c *gin.Context)
	GetWorkExperiences(c *gin.Context)
	CreateWorkExperience(c *gin.Context)
	CreateWorkExperiences(c *gin.Context)
	UpdateWorkExperience(c *gin.Context)
	DeleteWorkExperience(c *gin.Context)
	GetSkills(c *gin.Context)
	CreateSkill(c *gin.Context)
	CreateSkills(c *gin.Context)
	UpdateSkill(c *gin.Context)
	DeleteSkill(c *gin.Context)
}
