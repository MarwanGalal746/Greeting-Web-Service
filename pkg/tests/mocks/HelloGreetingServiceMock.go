package mocks

import (
	"Greeting-web-service/pkg/models"
	"github.com/stretchr/testify/mock"
)

type HelloServiceMock struct {
	mock.Mock
}

func (s *HelloServiceMock) Greet(person models.Person) (models.GreetingMsg, error) {
	s.Called(person)
	return models.GreetingMsg{Msg: "Hello Sherif!"}, nil
}
