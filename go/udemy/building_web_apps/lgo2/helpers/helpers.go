package helpers

import "time"

type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func (u *User) MyName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) MyAge() int {
	return u.Age
}
