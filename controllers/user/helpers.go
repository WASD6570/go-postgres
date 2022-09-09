package userController

import (
	"github.com/google/uuid"
	"github.com/wasd6570/go-postgres/db"
	"github.com/wasd6570/go-postgres/models"
)

func CreateUser(user *models.User) (*models.User, error) {

	user.ID = uuid.New()

	created_user := db.Conn.Create(&user)

	if created_user.Error != nil {
		return nil, created_user.Error
	}

	return user, nil
}
