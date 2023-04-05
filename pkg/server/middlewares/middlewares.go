package middlewares

import (
	"Greeting-web-service/pkg/errs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EnsureJSONContentType() gin.HandlerFunc {
	return func(c *gin.Context) {
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			c.AbortWithStatusJSON(http.StatusBadRequest, errs.ErrInvalidJson)
			return
		}
		c.Next()
	}
}

func EnsureRequestValidity() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method != http.MethodPost {
			c.AbortWithStatusJSON(http.StatusMethodNotAllowed, errs.ErrMethodNotAllowed)
			return
		}
		if c.Request.URL.Path != "/greet" {
			c.AbortWithStatusJSON(http.StatusNotFound, errs.ErrRouteNotFound)
			return
		}
		c.Next()
	}
}
