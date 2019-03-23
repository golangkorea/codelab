package handler

import (
	"encoding/gob"
	"text/template"

	"github.com/golang-korea/codelab/model"
	"github.com/gorilla/sessions"
)

var (
	store *sessions.CookieStore
	tmpl  *template.Template
)

func init() {
	store = sessions.NewCookieStore([]byte("secret"))
	tmpl = template.Must(template.ParseGlob("templates/*.html"))

	// Register User struct value type
	gob.Register(model.User{})
}