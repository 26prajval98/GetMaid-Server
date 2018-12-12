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
	"GoCar"
	"fmt"
	"github.com/rs/cors"
)

func main() {

	defer database.CloseDb()

	GoCar.HandlePath("/", index.Handler, jwt.VerifyJWT)
	GoCar.HandlePath("/signup", signup.Handler)
	GoCar.HandlePath("/login", local.Handler, middlewares.EnableCors)
	GoCar.HandlePath("/verify", verifyphone.Handler, jwt.VerifyJWT)

	//Common
	GoCar.HandlePath("/details", getdetails.Handler, jwt.VerifyJWT)
	GoCar.HandlePath("/ismaid", ismaid.Handler, jwt.VerifyJWT)
	GoCar.HandlePath("/pending", pending.Handler, jwt.VerifyJWT)
	GoCar.HandlePath("/all", all.Handler, jwt.VerifyJWT)

	// Maid Paths
	GoCar.HandlePath("/maidservices", maidservices.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	GoCar.HandlePath("/maidonline", maidonline.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	GoCar.HandlePath("/earnings", earnings.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	GoCar.HandlePath("/getname", methods.Handler, jwt.VerifyJWT, middlewares.IsMaid)
	GoCar.HandlePath("/maidid", maidid.Handler, jwt.VerifyJWT, middlewares.IsMaid)

	// Hirer Paths
	GoCar.HandlePath("/maidsearch", maidsearch.Handler, jwt.VerifyJWT, middlewares.IsHirer)
	GoCar.HandlePath("/done", done.Handler, jwt.VerifyJWT, middlewares.IsHirer)

	fmt.Println("Server Started")

	GoCar.SetHandler(cors.AllowAll().Handler)

	GoCar.StartServer(3000)
}
