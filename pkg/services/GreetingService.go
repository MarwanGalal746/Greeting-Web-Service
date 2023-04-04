package services

import "Greeting-web-service/pkg/models"

type GreetingService interface {
	Greet(msg models.Person) (models.GreetingMsg, error)
}
