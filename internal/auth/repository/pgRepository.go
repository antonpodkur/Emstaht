package repository

import (
	"errors"

	"github.com/antonpodkur/Emstaht/internal/auth"
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.Repository {
	return &authRepository{db: db}
}

func (r *authRepository) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) Update(user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) Delete(userID uuid.UUID) error {
	err := r.db.Delete(&models.User{}, userID).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *authRepository) GetAll() ([]*models.User, error) {
	var users []*models.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *authRepository) GetByID(userID uuid.UUID) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user, "id = ?", userID).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *authRepository) GetByName(name string) (*models.User, error) {
	var user *models.User
	err := r.db.Find(&user, "name = ?", name).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *authRepository) GetByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
