package googlelogin


import (
	"GetMaid/handlers/methods"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3000/google_login",
	ClientID:     "737486596598-ujfh6k53tnj4l4phqo4vc1ggobjb26bv.apps.googleusercontent.com",
	ClientSecret: "Ge0Ft6Ywy16w5gtyK9IaVFEr",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

type GoogleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Gender        string `json:"gender"`

}

func Handler(res http.ResponseWriter, req *http.Request) error {
	var e error

	defer methods.ErrorHandler(res, &e)

	switch {
	case methods.CheckCase("GET", "/login/google_login", req):
		IndexHandler(res,req)
	case methods.CheckCase("GET", "/login/google_login/loginpage", req):
		LoginHandler(res,req)

	default:
		panic(404)
	}

	return e
}

