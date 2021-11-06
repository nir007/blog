package handlers

import (
	tpl "html/template"
	"net/http"
)

const cookieNameId = "uuid"

// IndexAction renders index page
func Index(w http.ResponseWriter, r *http.Request) {
	t := tpl.Must(tpl.ParseFiles(
		"./public/dist/index.html",
	))
	errTpl := t.Execute(w, nil)

	if errTpl != nil {
		panic(errTpl)
	}
}