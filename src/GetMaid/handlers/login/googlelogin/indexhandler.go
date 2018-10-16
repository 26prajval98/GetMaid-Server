package googlelogin
import (
	"fmt"
	"net/http"
)

func indexHandler(res http.ResponseWriter, req *http.Request)error {
	fmt.Fprintln(res, "<a href='/login/google_login/loginpage'>Log in with Google</a>")
	return nil
}
