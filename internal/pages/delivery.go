package pages

import "github.com/gin-gonic/gin"

type Handlers interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	GetByUserId(c *gin.Context)
}
