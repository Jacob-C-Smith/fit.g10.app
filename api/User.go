// package declaration
package api

// imports
import (
	"fit/application"
	"html/template"
	"net/http"
)

// data
var (
	user_template *template.Template
)

func init() {

	// load the template from the file system
	user_template, _ = template.ParseFiles("template/main.tmpl")
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	user := application.UserHandler()

	// respond with the menu item HTML
	user_template.Execute(w, user)
}
