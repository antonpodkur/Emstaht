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
	err := r.db.Preload("User").Preload("WorkExperiences").Preload("Skills").First(&page, "user_id = ?", userId).Error
	if err != nil {
		return nil, err
	}
	return page, nil
}

func (r *pagesRepository) GetWorkExperiences() ([]*models.WorkExperience, error) {
	var exps []*models.WorkExperience
	if err := r.db.Find(&exps).Error; err != nil {
		return nil, err
	}
	return exps, nil
}

func (r *pagesRepository) CreateWorkExperience(exp *models.WorkExperience) (*models.WorkExperience, error) {
	err := r.db.Create(exp).Error
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func (r *pagesRepository) CreateWorkExperiences(exps []*models.WorkExperience) ([]*models.WorkExperience, error) {
	err := r.db.Create(exps).Error
	if err != nil {
		return nil, err
	}
	return exps, nil
}

func (r *pagesRepository) UpdateWorkExperience(exp *models.WorkExperience) (*models.WorkExperience, error) {
	err := r.db.Save(exp).Error
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func (r *pagesRepository) DeleteWorkExperience(id uuid.UUID) error {
	err := r.db.Delete(&models.WorkExperience{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *pagesRepository) GetSkills() ([]*models.Skill, error) {
	var skills []*models.Skill
	if err := r.db.Find(&skills).Error; err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *pagesRepository) CreateSkill(skill *models.Skill) (*models.Skill, error) {
	err := r.db.Create(skill).Error
	if err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *pagesRepository) CreateSkills(skills []*models.Skill) ([]*models.Skill, error) {
	err := r.db.Create(skills).Error
	if err != nil {
		return nil, err
	}
	return skills, nil
}

func (r *pagesRepository) UpdateSkill(skill *models.Skill) (*models.Skill, error) {
	err := r.db.Save(skill).Error
	if err != nil {
		return nil, err
	}
	return skill, nil
}

func (r *pagesRepository) DeleteSkill(id uuid.UUID) error {
	err := r.db.Delete(&models.Skill{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
