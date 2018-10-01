package routers

import (
	"github.com/Shahid4183/Golang/REST/interfaces"
	"github.com/labstack/echo"
)

// User - this structure will be used to access user controller via user interface
type User struct {
	userController interfaces.User
}

// New - this function is constructor of User structure
func (User) New(userController interfaces.User) User {
	return User{
		userController: userController,
	}
}

// Register - this function will be used to register routes in user router
// All user related routes will be registered here
func (router User) Register(group *echo.Group) {
	group.POST("signup", router.signup)
}

// signup - will create a new user in database
func (router User) signup(context echo.Context) error {
	return nil
}
