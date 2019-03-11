package handler

import (
	"text/template"

	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
	tmpl  *template.Template
)

func init() {
	store = sessions.NewCookieStore([]byte("secret"))
	tmpl = template.Must(template.ParseGlob("templates/*.html"))
}