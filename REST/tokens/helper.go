package tokens

import (
	"github.com/Shahid4183/Golang/REST/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func GetEmailFromJWT(context echo.Context) (string, error) {
	if context.Get("user") == nil {
		return "", errors.New("Invalid token")
	}
	token := context.Get("user").(*jwt.Token)

	if token != nil {
		claims := token.Claims.(*models.UserClaims)

		return claims.Email, nil
	} else {
		return "", errors.New("Invalid token")
	}
}
