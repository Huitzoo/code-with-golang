package person

import (
	"database/sql"
	"net/http"
	"app/email"
	//"strconv"
	"math/rand"
	"fmt"

	"app/rest"	
)

type Credentials struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

func InitToken() string {
    b := make([]byte, 8)
	rand.Read(b)
    return fmt.Sprintf("%x", b)
}

func SignUP(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	request, err := rest.Request(r, &Person{})
	fmt.Println(request)
	if err != nil {
		rest.Response(w, http.StatusInternalServerError, nil, "Invalidate data","")
		return
	}
	fmt.Println(request)
	p := request.(*Person)
	err = p.insertIntoPerson(db)
	if err != nil {
		rest.Response(w, http.StatusInternalServerError, nil, err.Error(),"")
		return
	}
	rest.Response(w, http.StatusOK, p, "",InitToken())
}

func SignIn(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	p := &Person{}
	request, err := rest.Request(r, &Credentials{})
	if err != nil {
		rest.Response(w, http.StatusInternalServerError, nil, "Invalidate data","")
		return
	}
	err = p.loginPerson(db, request.(*Credentials).Mail, request.(*Credentials).Password)
	if err != nil {
		rest.Response(w, http.StatusUnauthorized, nil, err.Error(),"")
		return
	}
	rest.Response(w, http.StatusOK, p, "",InitToken())
}

func DashBoard(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	request, e := rest.Request(r, &Person{})
	if e != nil {
		rest.Response(w, http.StatusBadRequest, nil, e.Error(),"")
		return
	}
	p := request.(*Person)
	person,err := p.getPersonInfo(db)
	if err != nil {
		rest.Response(w, http.StatusBadRequest, nil, err.Error(),"")
		return
	}
	rest.Response(w, http.StatusOK, person, "","")
}


func GoogleSignin(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	request, err := rest.Request(r, &Person{})
	if err != nil {
		rest.Response(w, http.StatusInternalServerError, nil, "Invalidate data","")
		return
	}
	p := request.(*Person)
	err = p.insertIntoPerson(db)
	rest.Response(w, http.StatusOK, p, "","")
}

func Forgot(w http.ResponseWriter, r *http.Request, db *sql.DB){
	request, err := rest.Request(r, &Person{})
	if err != nil {
		rest.Response(w, http.StatusInternalServerError, nil, "Invalidate data","")
		return
	}
	p := request.(*Person)
	exists := p.existsPerson(db)
	if exists {
		email := &email.Email{From:"meerkats.developers@gmail.com",Pass:"OsJoJuDaAn",To:p.Mail}
		email.Send("http://localhost:8005/restpass")
		rest.Response(w, http.StatusOK, nil, "Email sent","")
		return
	}
	rest.Response(w, http.StatusBadRequest, nil,"Email don't exists","")
}

func ResetPass(w http.ResponseWriter, r *http.Request, db *sql.DB){

}


func GetAllPeople(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	p := &Person{}
	persons, err := p.getAllPeople(db)
	if err != nil {
		rest.Response(w, http.StatusBadRequest, nil, err.Error(),"")
		return
	}
	rest.Response(w, http.StatusOK, persons, "","")
}
