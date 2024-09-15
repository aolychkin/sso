package jwt

import (
	"time"

	"github.com/aolychkin/sso/internal/domain/models"
	"github.com/golang-jwt/jwt"
)

// TODO: покрыть тестами функцию (Unit-тесты)
func NewToken(user models.User, app models.App, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(duration).Unix()
	claims["app_id"] = app.ID

	// TODO: убрать потом секреты отсюда
	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
