package auth

import (
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/google/uuid"
)

type Repository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(userID uuid.UUID) error
	GetAll() ([]*models.User, error)
	GetByID(userID uuid.UUID) (*models.User, error)
	GetByName(name string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
}