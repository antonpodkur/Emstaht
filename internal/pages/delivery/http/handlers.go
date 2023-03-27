package http

import (
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/antonpodkur/Emstaht/internal/pages"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type pagesHandlers struct {
	pagesUsecase pages.Usecase
}

func NewPagesHandlers(pagesUsecase pages.Usecase) pages.Handlers {
	return &pagesHandlers{
		pagesUsecase: pagesUsecase,
	}
}

func (h *pagesHandlers) Create(c *gin.Context) {
	page := &models.Page{}
	if err := c.ShouldBindJSON(page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.pagesUsecase.Create(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
}

func (h *pagesHandlers) Update(c *gin.Context) {
	page := &models.Page{}
	if err := c.ShouldBindJSON(page); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPage, err := h.pagesUsecase.Update(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, updatedPage)
}

func (h *pagesHandlers) GetByUserId(c *gin.Context) {
	userIdString := c.Param("userId")
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	page, err := h.pagesUsecase.GetByUserId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, page)
}
