package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/golang-korea/codelab/model"
	"github.com/golang-korea/codelab/oauth"
	"github.com/gorilla/sessions"
)

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func GoogleAuthCallback(w http.ResponseWriter, r *http.Request) {
	// Validate state value
	session, _ := store.Get(r, "session")
	state := session.Values["state"]
	defer func() {
		delete(session.Values, "state")
		session.Save(r, w)
	}()
	if state != r.FormValue("state") {
		http.Error(w, "invalid session state", http.StatusUnauthorized)
		return
	}

	// Do oauth2 process and request the user information
	token, err := oauth.GoogleOAuthConf.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	client := oauth.GoogleOAuthConf.Client(context.Background(), token)
	userInfoResp, err := client.Get(oauth.GoogleUserInfoAPI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if userInfoResp.StatusCode != http.StatusOK {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	// Read and parse user information
	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var authUser model.User
	json.Unmarshal(userInfo, &authUser)

	// Save the user information in session to reuse later
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400,
	}
	session.Values["user"] = authUser.Email
	session.Values["username"] = authUser.Name
	session.Values["picture"] = authUser.Picture
	session.Save(r, w)

	// Redirect to profile page
	http.Redirect(w, r, "/profile", http.StatusFound)
}
