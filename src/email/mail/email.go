package mail

import (
	"log"
	"os"

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
	mailGun(e *Email) (int, error)
	sendGrid(e *Email) (int, error)
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
func Send(el Emailer, e *Email) (int, error) {
	respCode, err := el.mailGun(e)
	if err != nil {
		log.Println(err)
		respCode, err = el.sendGrid(e)
		if err != nil {
			log.Println(err)
			return respCode, err
		}
	}
	return respCode, nil
}

func (el EmailInfoer) mailGun(e *Email) (int, error) {
	mg := mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"), os.Getenv("MG_PUBLIC_API_KEY"))
	message := mailgun.NewMessage(
		e.From,
		e.Subject,
		e.Body,
		e.To)
	_, _, err := mg.Send(message)
	statusCode := 202
	if err != nil {
		log.Fatal(err)
		statusCode = 400
	}
	return statusCode, nil
}

func (el EmailInfoer) sendGrid(e *Email) (int, error) {
	sg := sendgrid.NewSendClient(os.Getenv("SG_API_KEY"))
	from := helpers.NewEmail(e.From, e.From)
	subject := e.Subject
	to := helpers.NewEmail(e.To, e.To)
	plainTextContent := e.Body
	htmlContent := e.Body
	message := helpers.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	response, err := sg.Send(message)
	if err != nil {
		log.Println(err)
	}
	return response.StatusCode, nil
}
