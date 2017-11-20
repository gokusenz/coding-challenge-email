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

func (f FakeEmailer) mailGun(e *Email) (int, error) {
	if f.MailGun == "" {
		err := errors.New("Error")
		return 400, err
	}

	return 202, nil
}

func (f FakeEmailer) sendGrid(e *Email) (int, error) {
	if f.SendGrid == "" {
		err := errors.New("Error")
		return 400, err
	}

	return 202, nil
}

func TestSendEmailSuccessWithMailgun(t *testing.T) {
	expected := 0

	f := FakeEmailer{
		MailGun: "Success",
	}

	e := Email{
		From:    "test@gmail.com",
		To:      "test@gmail.com",
		Subject: "tester",
		Body:    "tester",
	}

	result, _ := Send(f, &e)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}

func TestSendEmailSuccessWithSendGrid(t *testing.T) {
	expected := 0

	f := FakeEmailer{
		SendGrid: "Success",
	}

	e := Email{
		From:    "test@gmail.com",
		To:      "test@gmail.com",
		Subject: "tester",
		Body:    "tester",
	}

	result, _ := Send(f, &e)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}

func TestSendEmailAllFail(t *testing.T) {
	expected := 4
	expectedMsg := "Emails failed in sending. The error message is as followed: Error"

	f := FakeEmailer{}
	e := Email{
		From:    "test@gmail.com",
		To:      "test@gmail.com",
		Subject: "tester",
		Body:    "tester",
	}

	result, respMsg := Send(f, &e)

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

	if !reflect.DeepEqual(expectedMsg, respMsg) {
		t.Fatalf("Expected %v but got %v", expectedMsg, respMsg)
	}
}
