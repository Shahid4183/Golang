package controllers

import (
	"errors"

	"github.com/Shahid4183/Golang/REST/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

// User - this struct is wrapper arround gorm database instance
// We will use this structure to identify on User model
// We are doing database operations (create, insert, update and delete)
type User struct {
	db *gorm.DB
}

// MakeUserController - we will use this function to assign gorm instance object
// to our controller and add wrapper around it
func MakeUserController(db *gorm.DB) User {
	return User{
		db: db,
	}
}

// CreateUser - this function creates new user in database
// it accepts an argument of type User and returns error if any
func (controller User) CreateUser(user *models.User) error {
	if err := controller.db.Create(user).Error; err != nil {
		log.Error(err)
		return errors.New("Error while creating new user")
	}
	return nil
}

// SaveUser - this function updates user details in database
// it accepts an argument of type User and returns error if any
func (controller User) SaveUser(user *models.User) error {
	if err := controller.db.Save(user).Error; err != nil {
		log.Error(err)
		return errors.New("Error while creating new user")
	}
	return nil
}
