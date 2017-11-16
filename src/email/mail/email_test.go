package mail

import (
	"errors"
	"reflect"
	"testing"
)

type FakeEmailer struct {
	MailGun  string
	SendGrid string
}

func (f FakeEmailer) mailGun(e *Email) (string, error) {
	if f.MailGun == "" {
		err := errors.New("Error")
		return "500", err
	}

	return "202", nil
}

func (f FakeEmailer) sendGrid(e *Email) (string, error) {
	if f.SendGrid == "" {
		err := errors.New("Error")
		return "500", err
	}

	return "202", nil
}

func TestSetEmail(t *testing.T) {

	expected := &Email{
		To:      "nattawut.ru@gmail.com",
		From:    "gokusen.regis@gmail.com",
		Subject: "Test",
		Body:    "Test Set",
		Cc:      "gokusen.regis@gmail.com",
		Bcc:     "gokusen.regis@gmail.com",
	}

	result := New(expected.To, expected.From, expected.Subject, expected.Body, expected.Cc, expected.Bcc)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}

func TestSendEmailSuccessWithMailgun(t *testing.T) {
	expected := "202"
	// expected := errors.New("Error")

	f := FakeEmailer{
		MailGun: "Success",
	}
	result, _ := Send(f, &Email{})

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}

func TestSendEmailSuccessWithSendGrid(t *testing.T) {
	expected := "202"

	f := FakeEmailer{
		SendGrid: "Success",
	}
	result, _ := Send(f, &Email{})

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}
