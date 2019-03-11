package handler

import (
	"net/http"

	"github.com/golang-korea/codelab/oauth"
	"github.com/gorilla/sessions"
)

func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options = &sessions.Options{
		Path:   "/auth",
		MaxAge: 300,
	}
	state := randToken()
	session.Values["state"] = state
	session.Save(r, w)
	tmpl.ExecuteTemplate(w, "login.html", oauth.GoogleLoginURL(state))
}

func RenderProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	if _, ok := session.Values["user"]; !ok {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	tmpl.ExecuteTemplate(w, "profile.html", map[string]string{
		"user": session.Values["user"].(string),
		"username": session.Values["username"].(string),
	})
}