package googlelogin

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"log"
	"net/http"
)

func CallbackHandler(w http.ResponseWriter, r *http.Request)error{
	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Fprintf(w, "Code exchange failed with error %s\n", err.Error())
		log.Fatal("Code exchange failed")
	}

	if !token.Valid() {
		fmt.Fprintln(w, "Retreived invalid token")
		log.Fatal("Retreived invalid token")
	}

	fmt.Fprintln(w, token.AccessToken)

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		log.Printf("Error getting user from token %s\n", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)

	var user *GoogleUser
	err = json.Unmarshal(contents, &user)
	if err != nil {
		log.Printf("Error unmarshaling Google user %s\n", err.Error())
		log.Fatal("Cannot unmarshal")
	}

	fmt.Fprintf(w, "Email: %s\nName: %s\nImage link: %s\n", user.Email, user.Name, user.Picture)
return nil
}
