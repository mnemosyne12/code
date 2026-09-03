package main

import (
	"encoding/json"
	"log"

	"github.com/mnemosyne12/code/go/udemy/bwa/lgo/helpers"
)

func main() {
	u1 := helpers.User{
		FirstName: "Trevor",
		LastName:  "Sawler",
		Age:       13,
	}

	log.Println(u1.PrintIntro())

	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshalled []helpers.Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// write JSON from a struct
	var personSlice []helpers.Person

	var m1 helpers.Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	var m2 helpers.Person
	m2.FirstName = "Diane"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	personSlice = append(personSlice, m1)
	personSlice = append(personSlice, m2)

	newJson, err := json.MarshalIndent(personSlice, "", "    ")
	if err != nil {
		log.Printf("Error creating JSON string from struct: %v", err)
	}

	log.Println(string(newJson))

}
