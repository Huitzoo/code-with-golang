package person

import (
	"database/sql"
	"errors"
	"fmt"
)

type Person struct {
	ID       int    `json:"id" db:"id"`
	Mail     string `json:"mail" db:"mail"`
	Password string `json:"password" db:"password"`
	Name     string `json:"name" db:"name"`
	Phone    string `json:"phone" db:"phone"`
	Address  string `json:"address" db:"address"`
}

func (p *Person) toString() string {
	return "mail=" + p.Mail
}
func (p *Person) getAllPeople(db *sql.DB) ([]Person, error) {
	persons := []Person{}
	consult, err := getAllPeopleRepository(db)
	if err != nil {
		return nil, err
	}
	for consult.Next() {
		var p Person
		if err := consult.Scan(&p.ID, &p.Mail, &p.Password); err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}
	return persons, nil
}

func (p *Person) loginPerson(db *sql.DB, mail, password string) error {
	consult := authenticationRepository(db, mail, password)
	err := consult.Scan(&p.ID, &p.Mail, &p.Name)
	if err != nil {
		return err
	}
	return nil
}


func (p *Person) existsPerson(db *sql.DB) bool {
	err, exists := getPersonByEmailRepository(db, p.Mail)
	if err != nil {
		return exists
	}
	return exists
}

func (p *Person) getPersonInfo(db *sql.DB) (Person,error){
	person := Person{}
	fmt.Println(p.Mail)
	consult := infoPersonRespository(db, p.Mail)
	if err := consult.Scan(&person.Mail,&person.Name,&person.Phone,&person.Address); err!=nil {
		return person, err
	}

	return person,nil
}
func (p *Person) insertIntoPerson(db *sql.DB) error {
	err, exists := getPersonByEmailRepository(db, p.Mail)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Email exists")
	}
	_, err = insertPersonRepository(db, p)
	if err != nil {
		return err
	}
	return nil
}
