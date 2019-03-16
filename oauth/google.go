package oauth

import (
	"github.com/golang-korea/codelab/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const GoogleUserInfoAPI = "https://www.googleapis.com/oauth2/v3/userinfo"

var GoogleOAuthConf *oauth2.Config

// NewGoogleOAuth builds a google oauth configuration
func init() {
	conf := config.Get().OAuth2["google"]
	GoogleOAuthConf = &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  conf.RedirectURL,
		Scopes:       conf.Scopes,
		Endpoint:     google.Endpoint,
	}
}

// GetLoginURL returns oauth login url with state
func GoogleLoginURL(state string) string {
	return GoogleOAuthConf.AuthCodeURL(state)
}
