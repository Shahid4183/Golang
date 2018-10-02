package tokens

import (
	"encoding/hex"
	"time"

	"github.com/Shahid4183/Golang/REST/models"
	"github.com/btcsuite/btcd/btcec"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var (
	lifeTimeAccessToken  int64  = 1800  // life time AccessToken: 30 min
	lifeTimeRefreshToken int64  = 21600 // life time RefreshToken: 6 hours
	accessTokenKey       []byte         // key for encrypting accessToken
	refreshTokenKey      []byte         // key for encrypting refreshToken
)

// Set new key for encrypting accessToken.
// During execution of the program, the key can be set only once.
func SetAccessTokenKey(key []byte) error {
	if len(accessTokenKey) != 0 {
		return errors.New("key already exists")
	}
	accessTokenKey = key
	return nil
}

// Set new key for encrypting refreshToken.
// During execution of the program, the key can be set only once.
func SetRefreshTokenKey(key []byte) error {
	if len(refreshTokenKey) != 0 {
		return errors.New("key already exists")
	}
	refreshTokenKey = key
	return nil
}

// Generates a new key.
// Used mechanism for generating private keys of btcd
func Generate() (string, error) {
	privKey, err := btcec.NewPrivateKey(btcec.S256())
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(privKey.D.Bytes()), nil
}

// Generates a new token for access for private api
// Inputs:
// email: mail address of the user
// Outputs:
// field #1: access token,
// field #2: time when the token will expire
// field #3: error
func CreateAccessToken(email string) (string, int64, error) {
	if len(accessTokenKey) == 0 {
		return "", 0, errors.New("key is null")
	}
	return create(email, accessTokenKey, lifeTimeAccessToken)
}

// Generates a new refresh token
// Inputs:
// email: mail address of the user
// Outputs:
// field #1: refresh token,
// field #2: time when the token will expire
// field #3: error
func CreateRefreshToken(email string) (string, int64, error) {
	if len(refreshTokenKey) == 0 {
		return "", 0, errors.New("key is null")
	}
	return create(email, refreshTokenKey, lifeTimeRefreshToken)
}

// Parse refresh token
// Inputs:
// token: encrypted token
// Outputs:
// field #1: decrypted token fields
// field #2: error
// field #3: this "true" value if token is valid
func ParseRefreshToken(token string) (*models.UserClaims, error, bool) {
	if len(refreshTokenKey) == 0 {
		return nil, errors.New("key is null"), false
	}
	return parse(token, refreshTokenKey)
}

func create(email string, key []byte, lifeTimeToken int64) (string, int64, error) {
	if len(key) == 0 {
		return "", 0, errors.New("key is null")
	}
	v := time.Now().Unix() + lifeTimeToken
	claims := models.UserClaims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: v,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	return ss, v, err
}

func parse(tokenString string, key []byte) (*models.UserClaims, error, bool) {
	if len(key) == 0 {
		return nil, errors.New("key is null"), false
	}

	token, err := jwt.ParseWithClaims(tokenString, &models.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err, false
	}

	if claims, ok := token.Claims.(*models.UserClaims); ok && token.Valid {
		return claims, nil, true
	} else {
		return nil, err, false
	}
}
