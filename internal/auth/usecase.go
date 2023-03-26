package auth

import (
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/google/uuid"
)

type Usecase interface {
	Register(user *models.User) (*models.User, error)
	Login(user *models.User) (*models.UserWithToken, error)
	Update(user *models.User) (*models.User, error)
	Delete(userID uuid.UUID) error
	GetByID(userID uuid.UUID) (*models.User, error)
	GetByName(name string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetUsers() ([]*models.User, error)
}
