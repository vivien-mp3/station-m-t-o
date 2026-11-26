package templates

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func LoadTemplates() {
	templates = template.Must(template.ParseGlob("templates/views/*.html"))
}

func Render(w http.ResponseWriter, status int, name string, data any) {
	buffer := bytes.Buffer{}

	err := templates.ExecuteTemplate(&buffer, name, data)
	if err != nil {
		log.Printf("Erreur template: %s (%v)", name, err)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)
	buffer.WriteTo(w)
}

type ErrorModel struct{
	Code int
	Message string
}

func RenderError(w http.ResponseWriter, status int, message string) {
	Render(w, status, "error", ErrorModel{Code: status, Message: message})
}