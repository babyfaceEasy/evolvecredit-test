package user

import "time"

type User struct {
	ID int
	Email string
	FirstName string
	LastName string
	DOB time.Time
}

func NewUser(id int, email string, firstName string, lastName string, dob time.Time) *User {
	return &User{ID: id, Email: email, FirstName: firstName, LastName: lastName, DOB: dob}
}