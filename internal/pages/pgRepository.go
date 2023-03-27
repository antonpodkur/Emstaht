package pages

import (
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	Create(page *models.Page) (*models.Page, error)
	Update(page *models.Page) (*models.Page, error)
	GetByUserId(userId uuid.UUID) (*models.Page, error)
}
