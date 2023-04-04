package handlers

import (
	"Greeting-web-service/pkg/errs"
	"Greeting-web-service/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Start(port string) {
	router := gin.Default()

	helloGreetingHandler := NewHelloHandler(&services.HelloService{})

	//endpoint
	router.POST("/greet", helloGreetingHandler.Greet)

	router.Use(func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, errs.ErrMethodNotAllowed)
		}
		if c.Request.URL.Path != "/greet" {
			c.AbortWithStatusJSON(http.StatusNotFound, errs.ErrRouteNotFound)
		}
		return
		c.Next()
	})

	router.Run(":" + port)

}
