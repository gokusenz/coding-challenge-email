package mail_test

import (
	"email/mail"
	"reflect"
	"testing"
)

func TestSetEmail(t *testing.T) {

	expected := &mail.Email{
		To:      "nattawut.ru@gmail.com",
		From:    "gokusen.regis@gmail.com",
		Subject: "Test",
		Body:    "Test Set",
		Cc:      "gokusen.regis@gmail.com",
		Bcc:     "gokusen.regis@gmail.com",
	}

	result := mail.New(expected.To, expected.From, expected.Subject, expected.Body, expected.Cc, expected.Bcc)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}

func TestSendEmailSuccessWithMailgun(t *testing.T) {

}
