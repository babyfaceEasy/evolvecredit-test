package dto

import (
	"evolvecredit-test/internal/user"
	"time"
)

type UserDTO struct {
	ID        int     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DOB       time.Time `json:"dob"`
}

func GetUserDTOFromUser(user *user.User) UserDTO {
	return UserDTO{
		ID: user.ID,
		Email: user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		DOB: user.DOB,
	}
}
