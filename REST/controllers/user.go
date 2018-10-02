package controllers

import (
	"errors"

	"github.com/Shahid4183/Golang/REST/models"
	"github.com/Shahid4183/Golang/REST/utility"
	"github.com/iancoleman/strcase"
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

// FindUser - this function finds user in database by name, email, phone or username
// it accepts two arguments and returns user (if found) and error (if any)
// searchBy - first argument accepts field name to search in
// searchString - second argument accepts string to search
func (controller User) FindUser(searchBy, searchString string) (*models.User, error) {
	// create new user object
	user := new(models.User)
	// create condition
	var condition string
	if utility.ReflectStructField(user, searchBy) == nil {
		condition = strcase.ToSnake(searchBy) + " ilike %" + searchString + "%"
	}
	// query database
	if err := controller.db.Where(condition).Find(user).Error; err != nil {
		log.Error(err)
		return nil, errors.New("User does not exist")
	}
	// if no error, return user
	return user, nil
}

// SaveUser - this function updates user details in database
// it accepts an argument of type User and returns error if any
func (controller User) SaveUser(user *models.User) error {
	if err := controller.db.Save(user).Error; err != nil {
		log.Error(err)
		return errors.New("Error while updating user details")
	}
	return nil
}

// RemoveUser - this function removes user details from database
// it accepts an argument of type User and returns error if any
func (controller User) RemoveUser(user *models.User) error {
	if err := controller.db.Delete(user).Error; err != nil {
		log.Error(err)
		return errors.New("Error while deleting user details")
	}
	return nil
}
