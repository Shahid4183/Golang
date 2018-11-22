package models

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// User - this model stores information about users
// We will use this model to auto migrate and automatically create table
// in database using gorm library
type User struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	ReturnedUser
	Password string `json:"password"`
}

// ReturnedUser - this struct will be used to return user information in API call
// which requires user information in response
type ReturnedUser struct {
	Username string `gorm:"max=64" json:"username" validate:"required,max=64"`
	FullName string `gorm:"max=50" json:"fullname" validate:"required,max=50"`
	Email    string `gorm:"unique" json:"email" validate:"required,email"`
	Mobile   string `gorm:"max=15,unique" json:"mobile" validate:"required,max=15"`
}

// UserClaims will hold claim in jwt
type UserClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
