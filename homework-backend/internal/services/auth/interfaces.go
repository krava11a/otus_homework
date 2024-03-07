package auth

import (
	"homework-backend/internal/models"
)

type UserCreater interface {
	CreateUser(user models.User) (string, error)
}

type UserProvider interface {
	GetUserById(user_id string) (user models.User, err error)
}

type AppProvider interface {
	App() (models.App, error)
}
