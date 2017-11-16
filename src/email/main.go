package main

import (
	"email/mail"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome!")
}

func handlerEmail(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	t1 := time.Now()
	decoder := json.NewDecoder(r.Body)
	var e *mail.Email
	err := decoder.Decode(&e)
	if err != nil {
		log.Println(err)
		rw.WriteHeader(400)
		rw.Write([]byte("Error"))
	} else {
		log.Println(e.Subject)
		el := mail.EmailInfoer{}
		resp, err := mail.Send(el, e)
		if err != nil {
			log.Println(err)
		}
		rw.WriteHeader(resp)
		rw.Write([]byte("Success"))
	}
	t2 := time.Now()
	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintf(os.Stderr, "Exception: %v\n", err)
			os.Exit(1)
		}
	}()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/email", handlerEmail)
	http.ListenAndServe(":8080", nil)

	// e := &mail.Email{
	// 	To:      "nattawut.ru@gmail.com",
	// 	From:    "gokusen.regis@gmail.com",
	// 	Subject: "Test",
	// 	Body:    "Test Set",
	// 	Cc:      "gokusen.regis@gmail.com",
	// 	Bcc:     "gokusen.regis@gmail.com",
	// }
	// el := mail.EmailInfoer{}
	// resp, _ := mail.Send(el, e)
	// fmt.Println(resp)

}

func errResponse(rw http.ResponseWriter) {
	rw.WriteHeader(500)
	rw.Write([]byte("Error"))
}
