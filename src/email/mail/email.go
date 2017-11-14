package mail

import (
	"fmt"
	"log"

	"gopkg.in/mailgun/mailgun-go.v1"
	sendgrid "gopkg.in/sendgrid/sendgrid-go.v2"
)

const mg = mailgun.NewMailgun(yourdomain, ApiKey, publicApiKey)
const sg = sendgrid.NewSendGridClient("sendgrid_user", "sendgrid_key")

// Email represents an email.
type Email struct {
	To      string
	From    string
	Subject string
	Body    string
	Cc      string
	Bcc     string
}

// New Email
func New() *Email {
	return &Email{}
}

// Set method
func (e *Email) Set(To string, From string, Subject string, Body string, Cc string, Bcc string) *Email {
	return &Email{
		To:      To,
		From:    From,
		Subject: Subject,
		Body:    Body,
		Cc:      Cc,
		Bcc:     Bcc,
	}
}

func (e *Email) Send() error {

	err := mailGun(e)
	if err != nil {
		log.Fatal(err)
		err = sendGrid(e)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func mailGun(e *Email) error {
	message := mailgun.NewMessage(
		e.From,
		e.Subject,
		e.Body,
		e.To)
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
}

func sendGrid(e *Email) error {

	message := sendgrid.NewMail()
	message.AddTo(e.To)
	message.SetSubject(e.Subject)
	message.SetHTML(e.Body)
	message.SetFrom(e.From)
	_, err := sg.Send(message)
	if err != nil {
		log.Fatal(err)
		return err
	}
}
