package handlers

import (
	"Greeting-web-service/pkg/errs"
	"Greeting-web-service/pkg/models"
	"Greeting-web-service/pkg/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloHandler struct {
	helloService services.GreetingService
}

func (h *HelloHandler) Greet(c *gin.Context) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, errs.ErrInvalidJson)
		return
	}
	greetingMessage, err := h.helloService.Greet(person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errs.ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, greetingMessage)
}

func NewHelloHandler(service services.GreetingService) GreetingHandler {
	return &HelloHandler{helloService: service}
}
