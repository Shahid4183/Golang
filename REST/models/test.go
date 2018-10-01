package models

// User - this model stores information about users
// We will use this model to auto migrate and automatically create table
// in database using gorm library
type User struct {
	ID uint
	ReturnedUser
	Password string `json:"password"`
}

// ReturnedUser - this struct will be used to return user information in API call
// which requires user information in response
type ReturnedUser struct {
	Username string `json:"username"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}
