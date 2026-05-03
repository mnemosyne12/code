package main

import (
	"log"
	"time"
)

// Capitalize the struct so that it's availabile to other packages
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (u *User) myName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) myAge() int {
	return u.Age
}

func main() {

	userJohn := User{
		FirstName: "John",
		LastName:  "Smith",
		Age:       26,
	}

	userMike := User{
		FirstName: "Mike",
		LastName:  "Hunt",
		Age:       30,
	}

	userMap := make(map[string]User)

	userMap["john"] = userJohn
	userMap["mike"] = userMike

	names := []string{userJohn.FirstName, userMike.FirstName}

	for name, user := range userMap {
		log.Println(name, user.myName(), user.myAge())
	}

	log.Println(names)

}
