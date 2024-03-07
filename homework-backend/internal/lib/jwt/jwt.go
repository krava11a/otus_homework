package jwt

import (
	"homework-backend/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewToken(user_id string, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Добавляем в токен всю необходимую информацию
	claims := token.Claims.(jwt.MapClaims)
	claims["uuid"] = user_id
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app"] = "otus_homework"

	// Подписываем токен, используя секретный ключ приложения
	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
