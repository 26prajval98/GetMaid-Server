package main

import (
	"GetMaid/database"
	"GetMaid/handlers/all"
	"GetMaid/handlers/authentication/jwt"
	"GetMaid/handlers/authentication/local"
	"GetMaid/handlers/authentication/verifyphone"
	"GetMaid/handlers/done"
	"GetMaid/handlers/earnings"
	"GetMaid/handlers/getdetails"
	"GetMaid/handlers/index"
	"GetMaid/handlers/ismaid"
	"GetMaid/handlers/maidid"
	"GetMaid/handlers/maidonline"
	"GetMaid/handlers/maidsearch"
	"GetMaid/handlers/maidservices"
	"GetMaid/handlers/methods"
	"GetMaid/handlers/middlewares"
	"GetMaid/handlers/pending"
	"GetMaid/handlers/signup"
	"fmt"
	"github.com/26prajval98/GoCar"
	"github.com/rs/cors"
)

func main() {

	defer database.CloseDb()

	gocar.HandlePath("/", index.Handler, jwt.VerifyJWT)
	gocar.HandlePath("/signup", signup.Handler)
	gocar.HandlePath("/login", local.Handler, middlewares.EnableCors)
	gocar.HandlePath("/verify", verifyphone.Handler, jwt.VerifyJWT)

	//Common
	gocar.HandlePath("/details", getdetails.Handler, jwt.VerifyJWT)
	gocar.HandlePath("/ismaid", ismaid.Handler, jwt.VerifyJWT)
	gocar.HandlePath("/pending", pending.Handler, jwt.VerifyJWT)
	gocar.HandlePath("/all", all.Handler, jwt.VerifyJWT)

	// Maid Paths
	gocar.HandlePath("/maidservices", maidservices.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	gocar.HandlePath("/maidonline", maidonline.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	gocar.HandlePath("/earnings", earnings.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	gocar.HandlePath("/getname", methods.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	gocar.HandlePath("/maidid", maidid.Handler, jwt.VerifyJWT, middlewares.IsMaid)

	// Hirer Paths
	gocar.HandlePath("/maidsearch", maidsearch.Handler, jwt.VerifyJWT, middlewares.IsHirer)
	gocar.HandlePath("/done", done.Handler, jwt.VerifyJWT, middlewares.IsHirer)

	fmt.Println("Server Started")

	gocar.SetHandler(cors.AllowAll().Handler)

	gocar.StartServer(3000)
}
