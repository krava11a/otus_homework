package jwt

import (
	"fmt"
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

func GetUserId(token string, app models.App) (string, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte([]byte(app.Secret)), nil
	})
	if err != nil {
		return "", err
	}

	// do something with decoded claims
	for key, val := range claims {
		if key == "uuid" {
			return fmt.Sprintf("%v", val), nil
		}
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	return "", nil
}
