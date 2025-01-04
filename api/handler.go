package api

import (
	"net/http"

	"github.com/EredinHawk/rest_api_authentication/auth"
	"github.com/gin-gonic/gin"
)

const quote = "Don't communicate by sharing memory, share memory by communicating."

// Quote - защищенный обработчик конечной точки GET localhost:8080/guote, который возвращает цитату.
// Обязательное требование - наличие заголовка Authorization, с валидным JWT токеном.
func Quote(c *gin.Context) {
	err := auth.TokenValidation(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "the authorization token is not valid", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"quote": quote})
}
