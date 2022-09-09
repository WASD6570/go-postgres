package authControllers

import (
	userControllers "github.com/wasd6570/go-postgres/controllers/user"
	"github.com/wasd6570/go-postgres/db"
	"github.com/wasd6570/go-postgres/models"
)

func FindOneAuth(auth *models.Auth) models.Auth {
	var Auth models.Auth

	db.Conn.Where(&models.Auth{Email: auth.Email, Password: auth.Password}).First(&Auth)

	return Auth

}

func CreateUserAndAuth(auth *models.Auth, user *models.User) *models.User {

	createdUser, err := userControllers.CreateUser(user)

	if err != nil {
		panic(err)
	}

	auth.UserID = createdUser.ID

	result := db.Conn.Create(&auth)

	if result.Error != nil {
		panic(result.Error)
	}
	return createdUser
}
