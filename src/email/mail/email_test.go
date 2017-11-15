package mail_test

import (
	"email/mail"
	"reflect"
	"testing"
)

func TestSetContent(t *testing.T) {

	expected := &mail.Email{
		To:      "nattawut.ru@gmail.com",
		From:    "gokusen.regis@gmail.com",
		Subject: "Test",
		Body:    "Test Set",
		Cc:      "gokusen.regis@gmail.com",
		Bcc:     "gokusen.regis@gmail.com",
	}

	e := mail.New()
	result := e.Set(expected.To, expected.From, expected.Subject, expected.Body, expected.Cc, expected.Bcc)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}
