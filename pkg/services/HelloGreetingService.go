package services

import (
	"Greeting-web-service/pkg/models"
	"fmt"
)

type HelloService struct {
}

const hello = "Hello %s!"

func (s *HelloService) Greet(person models.Person) (models.GreetingMsg, error) {
	var helloGreetingMsg models.GreetingMsg
	helloGreetingMsg.Msg = fmt.Sprintf(hello, person.Name)
	return helloGreetingMsg, nil
}
