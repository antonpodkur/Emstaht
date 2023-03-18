package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

type User struct {
	Base
	Name     string         `gorm:"size:255;not null;" json:"firstname"`
	Surname  string         `gorm:"size:255;not null;" json:"surname"`
	Email    string         `gorm:"size:255;not null;unique;" json:"email"`
	Dob      datatypes.Date `json:"dob"`
	Password string         `gorm:"not null;" json:"password"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePasswords(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
