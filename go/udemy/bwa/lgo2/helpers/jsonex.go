package helpers

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

var People []Person

func NewPerson(fname, lname, hcolor string, hasDog bool) Person {
	return Person{
		FirstName: fname,
		LastName:  lname,
		HairColor: hcolor,
		HasDog:    hasDog,
	}
}
