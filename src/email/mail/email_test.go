package mail_test

import (
	"email/mail"
	"reflect"
	"testing"
)

func TestSetContent(t *testing.T) {

	expected := &mail.Email{
		To:      "example@gmail.com",
		From:    "nattawut.ru@gmail.com",
		Subject: "Test",
		Body:    "Test Set",
		Cc:      "nattawut.cc@gmail.com",
		Bcc:     "nattawut.bcc@gmail.com",
	}

	e := mail.New()
	result := e.Set(expected.To, expected.From, expected.Subject, expected.Body, expected.Cc, expected.Bcc)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}
