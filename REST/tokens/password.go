package tokens

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	config "github.com/spf13/viper"
)

// Encrypt : utility function to encrypt given text using AES encryption
// Accepts text string to be encrypted
func Encrypt(text string) (string, error) {
	keyStr := config.GetString("passwordKey")
	keyBytes := sha256.Sum256([]byte(keyStr))
	key := keyBytes[:]

	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt : utility function to decrypt given text using AES encryption
// Accepts text string to be decrypted
func Decrypt(cryptoText string) (string, error) {
	keyStr := config.GetString("passwordKey")
	keyBytes := sha256.Sum256([]byte(keyStr))
	key := keyBytes[:]

	ciphertext, err := base64.StdEncoding.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext), nil
}

// ComparePasswords : compares two encrypted passwords
// if oldPassword == newPassword then return true
// else returns false
// in case of errors it returns false and error
func ComparePasswords(oldPasswordEnc, newPasswordEnc string) error {
	newPassword, err := Decrypt(newPasswordEnc)
	if err != nil {
		log.Error(err)
		errors.New("Error while creating new Password")
	}

	oldPassword, err := Decrypt(oldPasswordEnc)
	if err != nil {
		log.Error(err)
		errors.New("Error while creating new Password")
	}
	if strings.Compare(oldPassword, newPassword) == 0 {
		return errors.New("Old password and new password can not be same")
	}
	return nil
}
