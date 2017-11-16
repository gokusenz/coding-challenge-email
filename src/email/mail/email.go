package mail

import (
	"fmt"
	"log"
	"os"
	"strconv"

	sendgrid "github.com/sendgrid/sendgrid-go"
	helpers "github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/mailgun/mailgun-go.v1"
)

// Email represents an email.
type Email struct {
	To      string
	From    string
	Subject string
	Body    string
	Cc      string
	Bcc     string
}

type EmailInfoer struct{}

type Emailer interface {
	mailGun(e *Email) (string, error)
	sendGrid(e *Email) (string, error)
}

// New method
func New(To string, From string, Subject string, Body string, Cc string, Bcc string) *Email {
	return &Email{
		To:      To,
		From:    From,
		Subject: Subject,
		Body:    Body,
		Cc:      Cc,
		Bcc:     Bcc,
	}
}

// Send method
func Send(el Emailer, e *Email) (string, error) {
	log.Println("SEND")
	resp, err := el.mailGun(e)
	if err != nil {
		log.Println(err)
		resp, err = el.sendGrid(e)
		if err != nil {
			log.Println(err)
			return resp, err
		}
	}
	fmt.Println(resp)
	return resp, nil
}

func (el EmailInfoer) mailGun(e *Email) (string, error) {
	mg := mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"), os.Getenv("MG_PUBLIC_API_KEY"))
	message := mailgun.NewMessage(
		e.From,
		e.Subject,
		e.Body,
		e.To)
	resp, id, err := mg.Send(message)
	if err != nil {
		log.Fatal(err)
		return resp, err
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return resp, nil
}

func (el EmailInfoer) sendGrid(e *Email) (string, error) {
	sg := sendgrid.NewSendClient(os.Getenv("SG_API_KEY"))
	from := helpers.NewEmail(e.From, e.From)
	subject := e.Subject
	to := helpers.NewEmail(e.To, e.To)
	plainTextContent := e.Body
	htmlContent := e.Body
	message := helpers.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := sg.Send(message)
	i := strconv.Itoa(response.StatusCode)
	if err != nil {
		log.Println(err)

		return i, err
	}
	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)
	return i, nil
}
