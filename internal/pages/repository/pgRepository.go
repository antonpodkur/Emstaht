package repository

import (
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/antonpodkur/Emstaht/internal/pages"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type pagesRepository struct {
	db *gorm.DB
}

func NewPagesRepository(db *gorm.DB) pages.Repository {
	return &pagesRepository{
		db: db,
	}
}

func (r *pagesRepository) Create(page *models.Page) (*models.Page, error) {
	err := r.db.Create(page).Error
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (r *pagesRepository) Update(page *models.Page) (*models.Page, error) {
	err := r.db.Save(page).Error
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (r *pagesRepository) GetByUserId(userId uuid.UUID) (*models.Page, error) {
	var page *models.Page
	err := r.db.Preload("User").First(&page, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return page, nil
}
