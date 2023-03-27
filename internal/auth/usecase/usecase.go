package usecase

import (
	"errors"
	"github.com/antonpodkur/Emstaht/config"
	"github.com/antonpodkur/Emstaht/internal/auth"
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/antonpodkur/Emstaht/pkg/utils"
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
		return nil, errors.New("user with such email already exists")
	}
	if err != nil {
		return nil, err
	}

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
func (u *authUsecase) Login(user *models.User) (*models.UserWithToken, error) {
	foundUser, err := u.authRepository.GetByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if err := foundUser.ComparePasswords(user.Password); err != nil {
		return nil, err
	}

	foundUser.SanitizePassword()

	token, err := utils.GenerateJwtToken(foundUser, u.cfg)
	if err != nil {
		return nil, err
	}

	return &models.UserWithToken{
		User:  foundUser,
		Token: token,
	}, nil
}
func (u *authUsecase) Update(user *models.User) (*models.User, error) {
	return nil, nil
}
func (u *authUsecase) Delete(userID uuid.UUID) error {
	return nil
}
func (u *authUsecase) GetByID(userID uuid.UUID) (*models.User, error) {
	user, err := u.authRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	user.SanitizePassword()

	return user, nil
}
func (u *authUsecase) GetByName(name string) (*models.User, error) {
	return nil, nil
}
func (u *authUsecase) GetByEmail(email string) (*models.User, error) {
	return nil, nil
}
func (u *authUsecase) GetUsers() ([]*models.User, error) {
	return nil, nil
}
