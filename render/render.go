package render

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(
	template.ParseGlob("templates/*.html"),
)

func Render(w http.ResponseWriter, name string, data any) error {
	return tmpl.ExecuteTemplate(w, name, data)
}
