package views

import (
	"html/template"
	"log"
	"net/http"
)

/*Only server templates*/

func Home(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func Google(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/google.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

func DashBoard(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/register.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func Forgot(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/forgot.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func RestPass(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/restPass.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}