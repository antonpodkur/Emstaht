package http

import (
	"net/http"

	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/auth"
	httpModels "github.com/antonpodkur/Emstaht/internal/auth/delivery/dto"
	"github.com/gin-gonic/gin"
)

type authHandlers struct {
	cfg         *config.Config
	authUsecase auth.Usecase
}

func NewAuthHandlers(cfg *config.Config, authUsecase auth.Usecase) auth.Handlers {
	return &authHandlers{cfg: cfg, authUsecase: authUsecase}
}

// TODO: rewrite using user model with omitempty

// Register             godoc
// @Summary      Register new user
// @Description  Registers new user
// @Tags         auth
// @Accept		 json
// @Produce      json
// @Param 		userDto body dto.RegisterRequest true "RegisterJson"
// @Success      200
// @Router       /auth/register [post]
func (h *authHandlers) Register(c *gin.Context) {
	var request httpModels.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := request.MapRegisterRequestToUser()
	_, err := h.authUsecase.Register(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.Status(http.StatusCreated)
}
