package main

import (
	"log"

	"github.com/EredinHawk/rest_api_authentication/api"
	"github.com/EredinHawk/rest_api_authentication/auth"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/quote", 		 api.Quote)	  //Закрытая конечная точка API
	router.POST("/auth/sign-up", auth.SignUp) //Сервис аутентификации

	//Запуск локального сервера, который прослушивает входящие HTTP запросы по порту 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}