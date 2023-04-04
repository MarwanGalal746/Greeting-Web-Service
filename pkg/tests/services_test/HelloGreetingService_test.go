package services_test

import (
	"Greeting-web-service/pkg/models"
	"Greeting-web-service/pkg/services"
	"fmt"
	"testing"
)

func TestHelloService(t *testing.T) {
	// Create a new instance of the HelloService
	s := services.HelloService{}

	// Define test cases
	testCases := []struct {
		name    string
		person  models.Person
		want    models.GreetingMsg
		wantErr error
	}{
		{
			name: "valid input",
			person: models.Person{
				Name: "John",
			},
			want: models.GreetingMsg{
				Msg: fmt.Sprintf("Hello John!"),
			},
			wantErr: nil,
		},
		{
			name: "valid input 2",
			person: models.Person{
				Name: "",
			},
			want:    models.GreetingMsg{Msg: "Hello !"},
			wantErr: nil,
		},
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Greet method with the test input
			greeting, err := s.Greet(tc.person)

			// Check if the output matches the expected values
			if greeting.Msg != tc.want.Msg {
				t.Errorf("unexpected greeting message: got %q, want %q", greeting.Msg, tc.want.Msg)
			}
			if err != tc.wantErr {
				t.Errorf("unexpected error: got %v, want %v", err, tc.wantErr)
			}
		})
	}
}
