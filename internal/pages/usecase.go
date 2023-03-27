package pages

import (
	"github.com/antonpodkur/Emstaht/internal/models"
	"github.com/google/uuid"
)

type Usecase interface {
	Create(page *models.Page) (*models.Page, error)
	Update(page *models.Page) (*models.Page, error)
	GetByUserId(userId uuid.UUID) (*models.Page, error)
	GetWorkExperiences() ([]*models.WorkExperience, error)
	CreateWorkExperience(exp *models.WorkExperience) (*models.WorkExperience, error)
	CreateWorkExperiences(exps []*models.WorkExperience) ([]*models.WorkExperience, error)
	UpdateWorkExperience(exp *models.WorkExperience) (*models.WorkExperience, error)
	DeleteWorkExperience(id uuid.UUID) error
	GetSkills() ([]*models.Skill, error)
	CreateSkill(skill *models.Skill) (*models.Skill, error)
	CreateSkills(skills []*models.Skill) ([]*models.Skill, error)
	UpdateSkill(skill *models.Skill) (*models.Skill, error)
	DeleteSkill(id uuid.UUID) error
}
