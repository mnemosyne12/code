package helpers

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs(), "Legs")
}

// Dog Implements Animal interface
func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

// Gorilla implements Animal interface
func (g *Gorilla) Says() string {
	return "roar"
}

func (g *Gorilla) NumberOfLegs() int {
	return 4
}
