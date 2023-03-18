package dto

import (
	"time"

	"github.com/antonpodkur/Emstaht/internal/models"
	"gorm.io/datatypes"
)

type RegisterRequest struct {
	Name     string    `json:"name"`
	Surname  string    `json:"surname"`
	Email    string    `json:"email"`
	Dob      time.Time `json:"dob"`
	Password string    `json:"password"`
}

func (r *RegisterRequest) MapRegisterRequestToUser() models.User {
	user := models.User{}
	user.Name = r.Name
	user.Surname = r.Surname
	user.Email = r.Email
	user.Dob = datatypes.Date(r.Dob)
	user.Password = r.Password
	return user
}
