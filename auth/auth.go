package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SignUp - обработчик сервиса аутентификации, который регистрирует нового пользователья и возвращает JWT токен доступа
func SignUp(c *gin.Context) {
	//Принять логин и пароль
	credentails := Credentails{}
	if err := c.BindJSON(&credentails); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid request body, must be login : '...' and password: '...'"})
		return
	}

	//Проверить наличие в БД пользователя с таким логином
	if result := UserCheck(&credentails); result {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "the user with this username already exists"})
		return
	}

	//Если успех, то записать в БД нового пользователя
	FakeDataBase = append(FakeDataBase, credentails)

	//Вернуть access JWT токены
	jwt, err := GenirateTokens(&credentails)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "token generation error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "new user added", "JWT": jwt, "login": credentails.Login})
}

// NewUserCheck производит проверку на наличие пользователя в БД
func UserCheck(c *Credentails) bool {
	for _, v := range FakeDataBase {
		if v.Login == c.Login {
			return true
		}
	}
	return false
}
