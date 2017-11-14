package mail

// Message represents an email.
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

// SetContent method
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
