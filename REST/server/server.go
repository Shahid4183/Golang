package server

import (
	"net/http"

	"github.com/Shahid4183/Golang/REST/controllers"
	"github.com/Shahid4183/Golang/REST/routers"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Server handles http request response using echo framework
type Server struct {
	*echo.Echo
}

// New function creates new http server
func New(DB *gorm.DB, logger *lumberjack.Logger) (*Server, error) {
	server := Server{echo.New()}
	// Set middlewares to server
	// middleware to remove trailing slashes in api URI if any
	server.Pre(middleware.RemoveTrailingSlash())
	// Recover middleware recovers from panics anywhere in the chain,
	// prints stack trace and handles the control to the centralized http error handler
	server.Use(middleware.Recover())
	// middleware to enable Cross Origin Domains (CORS) functionality
	server.Use(middleware.CORS())
	// middleware to enable logging functionality
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: logger}))
	// Custome http error handler to handle errors returned from within the controllers and routers
	server.HTTPErrorHandler = func(err error, context echo.Context) {
		context.JSON(http.StatusInternalServerError, map[string]map[string]interface{}{
			"error": {
				"code":    500,
				"message": err.Error(),
			},
		})
	}
	// create controllers
	/* userController */
	userController := controllers.MakeUserController(DB)
	// create routers
	userRouter := routers.User{}.New(userController)
	// Register routes
	userRouter.Register(server.Group(""))
	return &server, nil
}
