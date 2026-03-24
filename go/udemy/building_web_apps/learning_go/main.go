package main

import (
	"github.com/mnemosyne12/code/go/udemy/building_web_apps/learning_go/helpers"
)

func main() {
	dog1 := helpers.Dog{
		Name:  "Jesus",
		Breed: "Corgi",
	}

	gorilla1 := helpers.Gorilla{
		Name:          "Christ",
		Color:         "Silver",
		NumberOfTeeth: 4,
	}

	helpers.PrintInfo(&dog1)

	helpers.PrintInfo(&gorilla1)
}
