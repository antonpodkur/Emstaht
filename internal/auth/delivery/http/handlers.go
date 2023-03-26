package http

import (
	"net/http"

	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/auth"
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/gin-gonic/gin"
)

type authHandlers struct {
	cfg         *config.Config
	authUsecase auth.Usecase
}

func NewAuthHandlers(cfg *config.Config, authUsecase auth.Usecase) auth.Handlers {
	return &authHandlers{cfg: cfg, authUsecase: authUsecase}
}

// Register             godoc
// @Summary      Register new user
// @Description  Registers new user
// @Tags         auth
// @Accept		 json
// @Produce      json
// @Param 		userDto body models.User true "RegisterJson"
// @Success      200
// @Router       /auth/register [post]
func (h *authHandlers) Register(c *gin.Context) {
	user := &models.User{}

	if err := c.ShouldBindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.authUsecase.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
}

func (h *authHandlers) Login(c *gin.Context) {
	type Login struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var req Login
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userWithToken, err := h.authUsecase.Login(&models.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, userWithToken)
}
