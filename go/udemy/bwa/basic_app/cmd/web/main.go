package main

import (
	"fmt"
	"net/http"

	"github.com/mnemosyne12/code/go/udemy/bwa/basic_app/pkg/handlers"
)

const portNumber = ":8081"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting the application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
