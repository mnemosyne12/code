package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// This function renders templates

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err2 := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err2 != nil {
		fmt.Println("Template failed to parse:", err2)
		return
	}
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}
