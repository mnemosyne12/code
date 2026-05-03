package main

import (
	"log"

	"github.com/mnemosyne12/code/go/udemy/building_web_apps/lgo2/helpers"
)

func main() {

	userJohn := helpers.User{
		FirstName: "John",
		LastName:  "Smith",
		Age:       26,
	}

	userMike := helpers.User{
		FirstName: "Mike",
		LastName:  "Hunt",
		Age:       30,
	}

	userMap := make(map[string]helpers.User)

	userMap["john"] = userJohn
	userMap["mike"] = userMike

	names := []string{userJohn.FirstName, userMike.FirstName}

	for name, user := range userMap {
		log.Println(name, user.MyName(), user.MyAge())
	}

	log.Println(names)

}
