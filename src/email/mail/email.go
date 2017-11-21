package mail

import (
	"errors"
	"log"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	helpers "github.com/sendgrid/sendgrid-go/helpers/mail"
	"gopkg.in/mailgun/mailgun-go.v1"
)

// Email represents an email.
type Email struct {
	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Cc      string `json:"cc"`
	Bcc     string `json:"bcc"`
}

type EmailInfoer struct{}

type Emailer interface {
	mailGun(e *Email) (int, error)
	sendGrid(e *Email) (int, error)
}

// Send : This is to send email by the email sender class and failover.
// Validate the from, to email address. Validate the subject and email content text.
// It will return an status code and message.
// Status  message
// 0       success
// 1       from email address invalid
// 2       no valid to email address
// 3       subject is empty
// 4       text is empty
// 5       email provider configuration not complete
// 6       all email sender failed
func Send(el Emailer, e *Email) (int, string) {

	if e.From == "" {
		return 1, "from email address invalid"
	}

	if e.To == "" {
		return 2, "To email address invalid"
	}

	if e.Subject == "" {
		return 3, "Subject invalid"
	}

	if e.Body == "" {
		return 4, "Body invalid"
	}

	respCode, err := el.sendGrid(e)
	if err != nil {
		log.Println("SendGrid failed. The error message is as followed: " + err.Error())
		// Failover to another email service provider.
		respCode, err = el.mailGun(e)
		if err != nil {
			log.Println("MailGun failed. The error message is as followed: " + err.Error())
		}
	}

	if respCode == 202 {
		return 0, "Success"
	} else if respCode == 500 {
		return 5, "Email provider configuration not complete"
	}
	return 6, "Emails failed in sending. The error message is as followed: " + err.Error()

}

// mailGun : This is the email sender using the MailGun implementation.
// return 202 if the sending success
// return 400 if the sending failed
// return 500 if email provider configuration not complete
func (el EmailInfoer) mailGun(e *Email) (int, error) {
	if os.Getenv("MG_DOMAIN") == "" || os.Getenv("MG_API_KEY") != "" || os.Getenv("MG_PUBLIC_API_KEY") != "" {
		return 500, errors.New("email provider configuration not complete")
	}
	mg := mailgun.NewMailgun(os.Getenv("MG_DOMAIN"), os.Getenv("MG_API_KEY"), os.Getenv("MG_PUBLIC_API_KEY"))
	message := mailgun.NewMessage(
		e.From,
		e.Subject,
		e.Body,
		e.To)
	_, _, err := mg.Send(message)
	statusCode := 202
	if err != nil {
		log.Println(err)
		statusCode = 400
	}
	return statusCode, err
}

// sendGrid : This is the email sender using the SendGrid implementation.
// return 202 if the sending success
// return 400 if the sending failed
// return 500 if email provider configuration not complete
func (el EmailInfoer) sendGrid(e *Email) (int, error) {
	if os.Getenv("SG_API_KEY") == "" {
		return 500, errors.New("email provider configuration not complete")
	}
	sg := sendgrid.NewSendClient(os.Getenv("SG_API_KEY"))
	from := helpers.NewEmail(e.From, e.From)
	subject := e.Subject
	to := helpers.NewEmail(e.To, e.To)
	plainTextContent := e.Body
	htmlContent := e.Body
	message := helpers.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	_, err := sg.Send(message)
	statusCode := 202
	if err != nil {
		log.Println(err)
		statusCode = 400
	}
	return statusCode, err
}
