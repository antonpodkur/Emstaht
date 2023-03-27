package usecase

import (
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/antonpodkur/Emstaht/internal/pages"
	"github.com/google/uuid"
)

type pagesUsecase struct {
	pagesRepository pages.Repository
}

func NewPagesUsecase(pagesRepository pages.Repository) pages.Usecase {
	return &pagesUsecase{
		pagesRepository: pagesRepository,
	}
}

func (u *pagesUsecase) Create(page *models.Page) (*models.Page, error) {
	createdPage, err := u.pagesRepository.Create(page)
	if err != nil {
		return nil, err
	}
	return createdPage, nil
}

func (u *pagesUsecase) Update(page *models.Page) (*models.Page, error) {
	updatedPage, err := u.pagesRepository.Update(page)
	if err != nil {
		return nil, err
	}
	return updatedPage, nil
}

func (u *pagesUsecase) GetByUserId(userId uuid.UUID) (*models.Page, error) {
	foundPage, err := u.pagesRepository.GetByUserId(userId)
	foundPage.User.SanitizePassword()
	if err != nil {
		return nil, err
	}
	return foundPage, nil
}
