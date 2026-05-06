package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8081"

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page") //So this writes back to the response
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(3, 4)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and the sum is: %d", sum))

}

func addValues(x, y int) int {
	return x + y
}
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
