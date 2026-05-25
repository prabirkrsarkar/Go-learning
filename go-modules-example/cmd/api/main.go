package main

import (
	"fmt"

	"github.com/prabirkrsarkar/Go-learning/go-modules-example/app/models"
)

// The type User is defined in the models package. We can use it here by importing the models package and then referring to it as models.User.
var user1 = models.User{
	Name: "Prabir",
}

func main() {
	//	user1.Name = "Prabir"
	fmt.Println("Hello, World from API!", user1.Name)
}
