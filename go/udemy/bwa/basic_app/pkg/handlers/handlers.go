package handlers

import (
	"fmt"
	"net/http"

	"github.com/mnemosyne12/code/go/udemy/bwa/basic_app/pkg/render"
)

// Home handler
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Home handler called.")
	render.RenderTemplate(w, "home.page.tmpl")
}

// About Handler
func About(w http.ResponseWriter, r *http.Request) {

	fmt.Println("About handler called.")
	render.RenderTemplate(w, "about.page.tmpl")

}
