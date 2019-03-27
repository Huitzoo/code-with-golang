package person

import (
	"database/sql"
	"fmt"
)

func getAllPeopleRepository(db *sql.DB) (*sql.Rows, error) {
	statement := fmt.Sprintf("select id,mail,password from person")
	return db.Query(statement)
}

/*
 -------------- For mysql database ----------------------------
func insertPersonRepository(db *sql.DB, data *Person) (sql.Result, error) {
	statement, err := db.Prepare("insert person set id=?,mail=?,password=?,name=?,phone=?,address=?")
	if err != nil {
		return nil, err
	}
	return statement.Exec(data.ID, data.Mail, data.Password, data.Name, data.Phone, data.Address)
}

func getPersonRepository(db *sql.DB, mail, password string) (*sql.Rows, error) {
	statement := fmt.Sprint("select mail,name from person where mail=? and password=?")
	return db.Query(statement, mail, password)
}

func getPersonByEmailRepository(db *sql.DB, mail string) (error, bool) {
	var exists bool
	statement := fmt.Sprint("select exists(select mail from person where mail=?)")
	return db.QueryRow(statement, mail).Scan(&exists), exists
}
*/

// This part i have to add for heroku database 
// -------------- for postgres databae -----------------------------
func insertPersonRepository(db *sql.DB, data *Person) (sql.Result, error) {
	statement := "insert into person(mail,password,name,phone,address) values($1,$2,$3,$4,$5)"
	return db.Exec(statement,data.Mail, data.Password, data.Name, data.Phone, data.Address)
}


func getPersonRepository(db *sql.DB, mail, password string) (*sql.Rows, error) {
	statement := fmt.Sprint("select mail,name from person where mail=$1 and password=$2")
	return db.Query(statement, mail, password)
}

func infoPersonRespository(db *sql.DB, mail string) *sql.Row{
	statement := fmt.Sprint("select mail,name,phone,address from person where mail=$1")
	return db.QueryRow(statement, mail)
}



func getPersonByEmailRepository(db *sql.DB, mail string) (error, bool) {
	var exists bool
	statement := "select exists(select mail from person where mail=$1)"
	return db.QueryRow(statement,mail).Scan(&exists), exists
}

func authenticationRepository(db *sql.DB, mail, password string) *sql.Row {
	//statement, err := db.Prepare("select mail, password from person where mail=? and password=?")
	statement := fmt.Sprint("select id,mail,name from person where mail=$1 and password=$2")
	return db.QueryRow(statement, mail, password)
}

