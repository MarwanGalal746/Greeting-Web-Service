package handlers_test

import (
	"Greeting-web-service/pkg/handlers"
	"Greeting-web-service/pkg/models"
	"Greeting-web-service/pkg/tests/mocks"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGreet(t *testing.T) {
	// Create a new mock Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Create a mock HelloService
	helloService := &mocks.HelloServiceMock{}

	helloHandler := handlers.NewHelloHandler(helloService)

	// Create a mock Person object
	person := models.Person{Name: "Sherif"}

	// Mock the service's Greet function
	expectedGreeting := models.GreetingMsg{Msg: "Hello Sherif!"}
	helloService.On("Greet", person).Return(expectedGreeting, nil)

	// Set the JSON request body in the Gin context
	requestBody := bytes.NewBufferString(`{
    "name" : "Sherif"
}`)
	c.Request, _ = http.NewRequest(http.MethodPost, "/greet", requestBody)
	c.Request.Header.Set("Content-Type", "application/json")

	helloHandler.Greet(c)

	// Verify the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Verify the response body
	expectedResponseBody := `{"message":"Hello Sherif!"}
`
	assert.Equal(t, expectedResponseBody, w.Body.String())

	// Verify that the mock service's Greet function was called
	helloService.AssertCalled(t, "Greet", person)
}
