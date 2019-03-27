package initapp

import (
	"os"
	"database/sql"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//app entities
	"app/person"
	
)
/*APP struct*/
/*
	Use App struct for define things to use, for example, Router, DB or another thing that you need.
	It's important init this Struct in web.go, because is like de brain of the app
*/
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, pass, dbName string) {
	var err error
	// connect to heroku database
	connectionString := os.Getenv("DATABASE_URL")
	//if connectionString is null we use the local database
	if connectionString == ""{
		//connectionString = fmt.Sprintf("%s:%s@/%s?charset=utf8", user, pass, dbName) // for local mysql 
		connectionString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",user,pass,dbName)
	}
	//in this case use postgres because heroku use postgres free, Sorry xD 
	a.DB, err = sql.Open("postgres",connectionString)
	//a.DB, err = sql.Open("mysql",connectionString) for mysql local
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Ok")
	}
	//init router
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	/*Define cors*/
	/*
		This part is important because, if the app work with an architecture of layers it have to accept	
		request for other domain and type of request
	*/
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "DELETE", "HEAD", "POST", "PUT", "OPTIONS"})
	fmt.Println("Run backend on port: ", addr)
	// run server with CORS
	err := http.ListenAndServe(addr, handlers.CORS(originsOk, headersOk, methodsOk)(a.Router))
	
	if err != nil {
		fmt.Println("I can't connect on a server: ", err)
	}
}

func (a *App) initializeRoutes() {

	/*Define all routers of app*/
	//Add db, beacuse I defined in de app struct and i need in the request of the app
	a.Router.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) { person.GetAllPeople(w, r, a.DB) }).Methods("GET")
	a.Router.HandleFunc("/singup", func(w http.ResponseWriter, r *http.Request) { person.SignUP(w, r, a.DB) }).Methods("POST")
	a.Router.HandleFunc("/singin", func(w http.ResponseWriter, r *http.Request) { person.SignIn(w, r, a.DB) }).Methods("POST", "GET")
	a.Router.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) { person.DashBoard(w, r, a.DB) }).Methods("GET", "POST")
	a.Router.HandleFunc("/googleSignin", func(w http.ResponseWriter, r *http.Request) { person.GoogleSignin(w, r, a.DB) }).Methods("POST")
	a.Router.HandleFunc("/forgot", func(w http.ResponseWriter, r *http.Request) { person.Forgot(w, r, a.DB) }).Methods("POST")
	a.Router.HandleFunc("/resetpass", func(w http.ResponseWriter, r *http.Request) { person.ResetPass(w, r, a.DB) }).Methods("POST")

}
