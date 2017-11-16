package mail

import (
	"errors"
	"reflect"
	"testing"
)

type FakeEmailer struct {
	To string
}

func (f FakeEmailer) mailGun(e *Email) (string, error) {
	if f.To == "" {
		err := errors.New("Error")
		return "", err
	}

	return "200", nil
}

func (f FakeEmailer) sendGrid(e *Email) (string, error) {
	if f.To == "" {
		err := errors.New("Error")
		return "500", err
	}

	return "200", nil
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
	expected := "200"
	// expected := errors.New("Error")

	f := FakeEmailer{
		To: "nattawut.ru@gmail.com",
	}
	result, _ := Send(f, &Email{})

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}
