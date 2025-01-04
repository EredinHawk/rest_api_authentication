package auth

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var secret = []byte("secret")

// GenirateTokens возвращает access токен аутентификации
func GenirateTokens(c *Credentails) (string, error) {
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Minute * 15).Unix(), //время жизни токена 15 минут
	})

	// Sign and get the complete encoded token as a string using the secret
	access_token, err := access.SignedString(secret)
	if err != nil {
		return "", err
	}
	return access_token, nil
}

// TokenValidation выполняет проверку токена JWT.
// Если такен не валидный, доступ к api будет отклонен.
func TokenValidation(c *gin.Context) error {
	//Извлекаем JWT из заголовка запроса
	token_string := c.GetHeader("Authorization")
	if token_string == "" {
		return fmt.Errorf("'Authorization' header is empty")
	}

	//Парсим токен
	token, err := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})
	if err != nil {
		return err
	}

	// Проверка токена на валидность
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return fmt.Errorf("Authorization token is not valid")
	}

	return nil
}