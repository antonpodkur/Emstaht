package models

import "gorm.io/datatypes"

type User struct {
	Base
	Name     string         `gorm:"size:255;not null;" json:"firstname"`
	Surname  string         `gorm:"size:255;not null;" json:"surname"`
	Email    string         `gorm:"size:255;not null;unique;" json:"email"`
	Dob      datatypes.Date `json:"dob"`
	Password string         `gorm:"not null;" json:"password"`
}
