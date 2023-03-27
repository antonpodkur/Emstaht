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

func (u *pagesUsecase) GetWorkExperiences() ([]*models.WorkExperience, error) {
	exps, err := u.pagesRepository.GetWorkExperiences()
	if err != nil {
		return nil, err
	}
	return exps, nil
}

func (u *pagesUsecase) CreateWorkExperience(exp *models.WorkExperience) (*models.WorkExperience, error) {
	createdExp, err := u.pagesRepository.CreateWorkExperience(exp)
	if err != nil {
		return nil, err
	}
	return createdExp, nil
}

func (u *pagesUsecase) CreateWorkExperiences(exps []*models.WorkExperience) ([]*models.WorkExperience, error) {
	createdExps, err := u.pagesRepository.CreateWorkExperiences(exps)
	if err != nil {
		return nil, err
	}
	return createdExps, nil
}

func (u *pagesUsecase) UpdateWorkExperience(exp *models.WorkExperience) (*models.WorkExperience, error) {
	updatedExp, err := u.pagesRepository.UpdateWorkExperience(exp)
	if err != nil {
		return nil, err
	}
	return updatedExp, nil
}

func (u *pagesUsecase) DeleteWorkExperience(id uuid.UUID) error {
	err := u.pagesRepository.DeleteWorkExperience(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *pagesUsecase) GetSkills() ([]*models.Skill, error) {
	skills, err := u.pagesRepository.GetSkills()
	if err != nil {
		return nil, err
	}
	return skills, nil
}

func (u *pagesUsecase) CreateSkill(skill *models.Skill) (*models.Skill, error) {
	createdSkill, err := u.pagesRepository.CreateSkill(skill)
	if err != nil {
		return nil, err
	}
	return createdSkill, nil
}

func (u *pagesUsecase) CreateSkills(skills []*models.Skill) ([]*models.Skill, error) {
	createdSkills, err := u.pagesRepository.CreateSkills(skills)
	if err != nil {
		return nil, err
	}
	return createdSkills, nil
}

func (u *pagesUsecase) UpdateSkill(skill *models.Skill) (*models.Skill, error) {
	updatedSkill, err := u.pagesRepository.UpdateSkill(skill)
	if err != nil {
		return nil, err
	}
	return updatedSkill, nil
}

func (u *pagesUsecase) DeleteSkill(id uuid.UUID) error {
	err := u.pagesRepository.DeleteSkill(id)
	if err != nil {
		return err
	}
	return nil
}
