package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"rabietf.me/go-assignment/models"
)

var (
	secretKey = []byte(os.Getenv("SECRET_TOKEN"))
)

// Service function that creates JWT for signed in user, using HS256 signing method.
// Returns "", error if something went wrong.
// Returns tokenString, nil if success.
func CreateToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(6 * time.Hour)

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	user.Password = ""

	claims["userID"] = user.ID
	claims["exp"] = expirationTime.Unix()

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
