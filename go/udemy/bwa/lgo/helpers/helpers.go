package helpers

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	DateTime  time.Time
}

func (u *User) PrintIntro() string {
	return fmt.Sprintf("Name: %s %s Age: %d", u.FirstName, u.LastName, u.Age)
}
