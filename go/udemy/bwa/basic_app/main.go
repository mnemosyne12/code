package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8081"

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page") //So this writes back to the response
	fmt.Println("Home handler called.")
}

// About Handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 4)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and the sum is: %d", sum))

	fmt.Println("About handler called.")

}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 10.0)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 10.0, f))

	fmt.Println("Divide handler called.")
}

func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}

	result := x / y
	return result, nil
}
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
