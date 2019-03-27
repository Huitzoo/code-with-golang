package initapp

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"app/views"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	fmt.Println("Run frontend on port: ", addr)
	err := http.ListenAndServe(addr, a.Router)
	if err != nil {
		fmt.Println("I can't connect on a server: ", err)
	}
}

func (a *App) initializeRoutes() {
	/*URL Home*/
	a.Router.HandleFunc("/", views.Home).Methods("GET")
	a.Router.HandleFunc("/dashboard", views.DashBoard).Methods("GET")
	a.Router.HandleFunc("/signup", views.Register).Methods("GET")
	a.Router.HandleFunc("/forgot", views.Forgot).Methods("GET")
	a.Router.HandleFunc("/restpass", views.Forgot).Methods("GET")
	a.Router.HandleFunc("/google", views.Google).Methods("GET")

	/*Server static failes */
	a.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
}
