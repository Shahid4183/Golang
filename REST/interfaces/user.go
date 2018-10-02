package interfaces

import "github.com/Shahid4183/Golang/REST/models"

// User - interface to provide access to user controller in simple and clean way
type User interface {
	CreateUser(user *models.User) error
	FindUser(searchBy, searchString string) (*models.User, error)
	SaveUser(user *models.User) error
	RemoveUser(user *models.User) error
}
