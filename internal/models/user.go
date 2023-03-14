package models

import "gorm.io/datatypes"

type User struct {
	Base
	Name     string         `json:"firstname"`
	Surname  string         `json:"surname"`
	Dob      datatypes.Date `json:"dob"`
	Password string         `json:"password"`
}
