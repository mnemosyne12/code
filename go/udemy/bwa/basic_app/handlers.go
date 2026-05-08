package main

import (
	"fmt"
	"net/http"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Home handler called.")
	renderTemplate(w, "home.page.tmpl")
}

// About Handler
func About(w http.ResponseWriter, r *http.Request) {

	fmt.Println("About handler called.")
	renderTemplate(w, "about.page.tmpl")

}
