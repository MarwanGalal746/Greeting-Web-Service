package models

type GreetingMsg struct {
	Msg string `json:"message"` // Name of the person is required in the json and should be consisted of alphabets only
}
