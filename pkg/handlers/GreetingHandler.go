package handlers

import (
	"github.com/gin-gonic/gin"
)

type GreetingHandler interface {
	Greet(c *gin.Context)
}
