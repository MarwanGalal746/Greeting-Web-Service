package models

type Person struct {
	Name string `json:"name" binding:"required,alpha" ` // Name of the person is required in the json and should be consisted of alphabets only
}
