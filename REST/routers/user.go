package routers

import (
	"net/http"

	"github.com/Shahid4183/Golang/REST/tokens"

	"github.com/Shahid4183/Golang/REST/interfaces"
	"github.com/Shahid4183/Golang/REST/models"
	Errors "github.com/Shahid4183/Golang/REST/utility/Errors"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	validator "gopkg.in/go-playground/validator.v9"
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
	group.POST("/signup", router.signup)
}

func (router User) getUser(context echo.Context) (*models.User, error) {
	// get email form email
	email, err := tokens.GetEmailFromJWT(context)
	if err != nil {
		return nil, Errors.InvalidJWTToken
	}
	// find user in database by email
	return router.userController.FindUser("Email", email)
}

// signup - will create a new user in database
func (router User) signup(context echo.Context) error {
	// Request body structure
	type Request struct {
		models.ReturnedUser        // all fileds of models.ReturnedUser will be available in Request structure
		Password            string `json:"password" validate:"required"`
	}
	// create request object
	request := new(Request)
	// bind request to context
	// this step takes request body send in api call and binds
	// request body parameters to request object
	if err := context.Bind(request); err != nil {
		log.Error(err)
		return Errors.InvalidRequest
	}
	// do server-side validation
	if err := validator.New().Struct(request); err != nil {
		log.Error(err)
		return Errors.ValidationError
	}
	// create a user
	user := new(models.User)
	// assign values to user
	user.Username = request.Username
	user.FullName = request.FullName
	user.Email = request.Email
	user.Mobile = request.Mobile
	user.Password = request.Password
	// create new user in database
	if err := router.userController.CreateUser(user); err != nil {
		return err
	}
	return context.JSON(http.StatusOK, "Registration successfull")
}
