package googlelogin

import (
	"github.com/dchest/uniuri"
	"net/http"
)

func LoginHandler(res http.ResponseWriter, req *http.Request)error {
	oauthStateString := uniuri.New()
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(res, req, url, http.StatusTemporaryRedirect)
	return nil
}
