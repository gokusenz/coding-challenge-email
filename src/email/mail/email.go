package mail

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	helpers "github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/mailgun/mailgun-go.v1"
)

var mg = mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"), os.Getenv("MG_PUBLIC_API_KEY"))
var sg = sendgrid.NewSendClient(os.Getenv("SG_API_KEY"))

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

	resp, err := mailGun(e)
	if err != nil {
		log.Fatal(err)
		resp, err = sendGrid(e)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println(resp)
	return nil
}

func mailGun(e *Email) (string, error) {
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

func sendGrid(e *Email) (string, error) {

	from := helpers.NewEmail(e.From, e.From)
	subject := e.Subject
	to := helpers.NewEmail(e.To, e.To)
	plainTextContent := e.Body
	htmlContent := ""
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
