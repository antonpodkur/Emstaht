package models

import "github.com/google/uuid"

type Page struct {
	Base
	About  string    `json:"about"`
	UserID uuid.UUID `gorm:"foreignKey:UserID;references:users(id)" json:"user_id"`
	User   User      `json:"user,omitempty" binding:"-"`
}
