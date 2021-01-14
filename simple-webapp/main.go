package main

import (
	"fmt"
	"simple-webapp/models"
)

func main() {
	u := models.User{
		ID:        1,
		FirstName: "Hung",
		LastName:  "Nguyen",
	}
	fmt.Println(u)
}
