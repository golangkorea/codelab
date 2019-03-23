package handler

import (
	"net/http"

	"github.com/golang-korea/codelab/model"
	"github.com/golang-korea/codelab/oauth"
	"github.com/gorilla/sessions"
)

func RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "auth")
	session.Options = &sessions.Options{
		Path:   "/auth",
		MaxAge: 300,
	}
	state := randToken()
	session.Values["state"] = state
	session.Save(r, w)
	tmpl.ExecuteTemplate(w, "login.html", oauth.GoogleAuthorizationURL(state))
}

func RenderProfile(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "auth")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, ok := session.Values["user"]; !ok {
		http.Error(w, "page not found", http.StatusNotFound)
		return
	}
	user := session.Values["user"].(model.User)
	tmpl.ExecuteTemplate(w, "profile.html", map[string]string{
		"name":  user.Name,
		"email": user.Email,
		"picture": user.Picture,
	})
}