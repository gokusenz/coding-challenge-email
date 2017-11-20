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
// 3       subject or text both are empty
// 4       all email sender failed
func Send(el Emailer, e *Email) (int, string) {

	if e.From == "" {
		return 1, "from email address invalid"
	}

	if e.To == "" {
		return 2, "To email address invalid"
	}

	if e.Subject == "" || e.Body == "" {
		return 3, "Subject or text invalid"
	}

	respCode, err := el.mailGun(e)
	if err != nil {
		log.Println(err)
		// Failover to another email service provider.
		respCode, err = el.sendGrid(e)
		if err != nil {
			log.Println(err)
		}
	}

	if respCode == 202 {
		return 0, "Success"
	}
	return 4, "Emails failed in sending. The error message is as followed: " + err.Error()

}

// mailGun : This is the email sender using the MailGun implementation.
// return 202 if the sending success
// return 400 if the sending failed
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

// sendGrid : This is the email sender using the SendGrid implementation.
// return 202 if the sending success
// return 400 if the sending failed
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
