package handler

import (
	"encoding/gob"
	"text/template"

	"github.com/golangkorea/codelab/model"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
	tmpl  *template.Template
)

func init() {
	store = sessions.NewCookieStore([]byte("secret-key"))
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	// Register User struct value type
	gob.Register(model.User{})
}