package models

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Page struct {
	Base
	About           string           `json:"about"`
	WorkExperiences []WorkExperience `json:"work_experiences" binding:"-"`
	Skills          []Skill          `json:"skills" binding:"-"`
	UserId          uuid.UUID        `gorm:"foreignKey:UserId;references:users(id)" json:"user_id" binding:"required"`
	User            User             `json:"user,omitempty" binding:"-"`
}

type WorkExperience struct {
	Base
	Name        string         `json:"name" binding:"required"`
	Descryption string         `json:"descryption" binding:"required"`
	StartDate   datatypes.Date `json:"start_date,omitempty"`
	EndDate     datatypes.Date `json:"end_date,omitempty"`
	PageId      uuid.UUID      `gorm:"foreignKey:PageId;references:pages(id)" json:"page_id" binding:"required"`
}

type Skill struct {
	Base
	Name        string    `json:"name" binding:"required"`
	Proficiency int       `json:"proficiency,omitempty"`
	PageId      uuid.UUID `gorm:"foreignKey:PageId;references:pages(id)" json:"page_id" binding:"required"`
}
