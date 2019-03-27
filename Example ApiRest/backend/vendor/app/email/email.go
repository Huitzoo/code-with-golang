package email

import (
    "net/smtp"
)

type Email struct{
	From string
	Pass string
	To string
} 


func(e *Email)Send(link string) error{

    from := smtp.PlainAuth("", e.From, e.Pass, "smtp.gmail.com")
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		from,
		e.From, 
		[]string{e.To}, []byte(link))
	if err != nil {
		return err
	}
	return nil
}
