package main

import (
	"encoding/json"
	"log"

	"github.com/devfeel/mapper"
)

type (
	// UserData has 2 field
	UserData struct {
		ID   int
		Name string
	}

	// UserResponse has 3 field
	UserResponse struct {
		ID   int
		Name string
		Note string
	}
)

func main() {
	// Initiate model
	userData := &UserData{
		ID:   1,
		Name: "John",
	}
	userResponse := &UserResponse{}

	// Map userData -> userResponse
	err := mapper.Mapper(userData, userResponse)
	if err != nil {
		log.Println(err)
	}
	res, _ := json.Marshal(userResponse)
	log.Println(string(res))
	
	// Update userResponse
	userResponse.ID = 2
	userResponse.Name = "Doe"
	userResponse.Note = "Hello"

	// Map userResponse -> userData
	err = mapper.Mapper(userResponse, userData)
	if err != nil {
		log.Println(err)
	}
	res, _ = json.Marshal(userData)
	log.Println(string(res))
}
