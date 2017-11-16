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
	expected := 202

	f := FakeEmailer{
		MailGun: "Success",
	}
	result, _ := Send(f, &Email{})

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}

func TestSendEmailSuccessWithSendGrid(t *testing.T) {
	expected := 202

	f := FakeEmailer{
		SendGrid: "Success",
	}
	result, _ := Send(f, &Email{})

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

}

func TestSendEmailAllFail(t *testing.T) {
	expected := 400
	expectedErr := errors.New("Error")

	f := FakeEmailer{}
	result, err := Send(f, &Email{})

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("Expected %v but got %v", expected, result)
	}

	if !reflect.DeepEqual(expectedErr, err) {
		t.Fatalf("Expected %v but got %v", expectedErr, err)
	}
}
