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

func (h *pagesHandlers) GetWorkExperiences(c *gin.Context) {
	exps, err := h.pagesUsecase.GetWorkExperiences()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, exps)
}

func (h *pagesHandlers) CreateWorkExperience(c *gin.Context) {
	exp := &models.WorkExperience{}
	if err := c.ShouldBindJSON(exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdExp, err := h.pagesUsecase.CreateWorkExperience(exp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdExp)
}

func (h *pagesHandlers) CreateWorkExperiences(c *gin.Context) {
	var exps []*models.WorkExperience
	if err := c.ShouldBindJSON(&exps); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdExps, err := h.pagesUsecase.CreateWorkExperiences(exps)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdExps)
}

func (h *pagesHandlers) UpdateWorkExperience(c *gin.Context) {
	exp := &models.WorkExperience{}
	if err := c.ShouldBindJSON(exp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedExp, err := h.pagesUsecase.UpdateWorkExperience(exp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedExp)
}

func (h *pagesHandlers) DeleteWorkExperience(c *gin.Context) {
	idString := c.Param("expId")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.pagesUsecase.DeleteWorkExperience(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}

func (h *pagesHandlers) GetSkills(c *gin.Context) {
	skills, err := h.pagesUsecase.GetSkills()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, skills)
}

func (h *pagesHandlers) CreateSkill(c *gin.Context) {
	skill := &models.Skill{}
	if err := c.ShouldBindJSON(skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdSkill, err := h.pagesUsecase.CreateSkill(skill)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdSkill)
}

func (h *pagesHandlers) CreateSkills(c *gin.Context) {
	var skills []*models.Skill
	if err := c.ShouldBindJSON(&skills); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdSkills, err := h.pagesUsecase.CreateSkills(skills)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdSkills)
}

func (h *pagesHandlers) UpdateSkill(c *gin.Context) {
	skill := &models.Skill{}
	if err := c.ShouldBindJSON(skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedSkill, err := h.pagesUsecase.UpdateSkill(skill)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSkill)
}

func (h *pagesHandlers) DeleteSkill(c *gin.Context) {
	idString := c.Param("skillId")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	err = h.pagesUsecase.DeleteSkill(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusOK)
}
