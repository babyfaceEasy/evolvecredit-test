package user

import (
	"time"

	"github.com/uptrace/bun"
)

type UserModel struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int `bun:"id,pk"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	FirstName     string
	LastName      string
	Email         string
	DOB           time.Time
}
