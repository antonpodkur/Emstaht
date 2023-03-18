package usecase

import (
	"errors"
	"time"

	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/auth"
	models "github.com/antonpodkur/Emstaht/internal/models"
	"github.com/google/uuid"
)

type authUsecase struct {
	cfg            *config.Config
	authRepository auth.Repository
}

func NewAuthUsecase(cfg *config.Config, authRepository auth.Repository) auth.Usecase {
	return &authUsecase{cfg: cfg, authRepository: authRepository}
}

func (u *authUsecase) Register(user *models.User) (*models.User, error) {
	existsUser, err := u.authRepository.GetByEmail(user.Email)
	if existsUser != nil {
		return nil, errors.New("User with such email already exists")
	}
	if err != nil {
		return nil, err
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = user.HashPassword()
	if err != nil {
		return nil, err
	}

	createdUser, err := u.authRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}
func (u *authUsecase) Login(user *models.User) (*models.User, error) {
	return nil, nil
}
func (u *authUsecase) Update(user *models.User) (*models.User, error){
	return nil, nil
}
func (u *authUsecase) Delete(userID uuid.UUID) error{
	return nil
}
func (u *authUsecase) GetByID(userID uuid.UUID) (*models.User, error){
	return nil, nil
}
func (u *authUsecase) GetByName(name string) (*models.User, error){
	return nil, nil
}
func (u *authUsecase) GetByEmail(email string) (*models.User, error){
	return nil, nil
}
func (u *authUsecase) GetUsers() ([]*models.User, error){
	return nil, nil
}
