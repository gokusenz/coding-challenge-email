package main

import (
	"email/mail"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func responseEmail(rw http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)
	var e *mail.Email
	err := decoder.Decode(&e)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400)
	} else {
		log.Println(e.Subject)
		el := mail.EmailInfoer{}
		resp, err := mail.Send(el, e)
		if err != nil {
			log.Println(err)
		}
		rw.WriteHeader(resp)
	}
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// }()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// http.HandleFunc("/email", responseEmail)
	// http.ListenAndServe(":80", nil)

	e := &mail.Email{
		To:      "nattawut.ru@gmail.com",
		From:    "gokusen.regis@gmail.com",
		Subject: "Test",
		Body:    "Test Set",
		Cc:      "gokusen.regis@gmail.com",
		Bcc:     "gokusen.regis@gmail.com",
	}
	el := mail.EmailInfoer{}
	resp, _ := mail.Send(el, e)

	fmt.Println(resp)

}

func errResponse(rw http.ResponseWriter) {
	rw.WriteHeader(500)
	rw.Write([]byte("Error"))
}
