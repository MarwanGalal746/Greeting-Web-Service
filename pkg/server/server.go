package server

import (
	"Greeting-web-service/pkg/handlers"
	"Greeting-web-service/pkg/server/middlewares"
	"Greeting-web-service/pkg/services"
	"github.com/gin-gonic/gin"
	"log"
)

func Start(port string) {
	router := gin.Default()

	helloGreetingHandler := handlers.NewHelloHandler(&services.HelloService{})

	router.Use(middlewares.EnsureJSONContentType())

	//endpoint
	router.POST("/greet", helloGreetingHandler.Greet)

	router.Use(middlewares.EnsureRequestValidity())

	err := router.Run(":" + port)
	if err != nil {
		log.Print("Invalid port number")
		return
	}

}
