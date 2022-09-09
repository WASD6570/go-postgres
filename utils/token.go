package apiUtils

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/wasd6570/go-postgres/models"
)

func EncodeToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func DecodeToken(tokenString string) (models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return models.User{}, err
	}
	claims := token.Claims.(jwt.MapClaims)
	user := models.User{
		ID: claims["user_id"].(uuid.UUID),
	}
	return user, nil
}
