package main

import (
	"encoding/json"
	"log"

	"github.com/mnemosyne12/code/go/udemy/building_web_apps/lgo2/helpers"
)

func main() {
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

	//write json from a struct
	flash := helpers.NewPerson("Wally", "West", "red", false)
	wonder_woman := helpers.NewPerson("Diana", "Prince", "black", false)

	helpers.People = append(helpers.People, flash)
	helpers.People = append(helpers.People, wonder_woman)

	newJson, err := json.MarshalIndent(helpers.People, "", "    ")

	log.Println("People JSON:", string(newJson))

}
