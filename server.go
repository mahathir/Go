package main

import (
	// Standard library packages

	"net/http"

	mgo "gopkg.in/mgo.v2"

	// Internal dependencies
	"three_tiers/api/controllers"
	"three_tiers/business"
	"three_tiers/business/interfacebusiness"
	"three_tiers/data_access/interfaceda"
	"three_tiers/data_access/mongodb"
	"three_tiers/utilities"

	// Third party packages
	"github.com/julienschmidt/httprouter"
)

var appCfg = utilities.NewAppConfig()
var diConst = utilities.NewDIConstants()
var iocCont = utilities.NewIocContainer()

func main() {
	// Initializing IoC
	initIoc()

	// Instantiate a new router
	r := httprouter.New()

	// Initializing User Controller and its route
	initUserContr(r)

	// Fire up the server
	http.ListenAndServe("localhost:3000", authMiddleware(r))
}

func initUserContr(r *httprouter.Router) {
	uc := &controllers.UserController{UserService: iocCont.Retrieve(diConst.IUsrServ).(interfacebusiness.IUserService)}

	// Get a user resource
	r.GET("/user/:id", uc.GetUser)
}

func initIoc() {
	// DataAccess
	session := getSession(appCfg.MongoDbConn)
	iocCont.Reg(diConst.IUsrDa, &mongodb.UserDataAccess{Session: session})

	// Business Services
	iocCont.Reg(diConst.IUsrServ, &business.UserService{Da: iocCont.Retrieve(diConst.IUsrDa).(interfaceda.IUserDataAccess)})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession(mongoconn string) *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial(mongoconn)

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}
