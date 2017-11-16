package main

import (
	"email/mail"
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func responseEmail(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Write([]byte("Email"))
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

	e := mail.New("nattawut.ru@gmail.com", "gokusen.regis@gmail.com", "Test", "Test Set", "gokusen.regis@gmail.com", "gokusen.regis@gmail.com")
	el := mail.EmailInfoer{}
	resp, _ := mail.Send(el, e)

	fmt.Println(resp)

}

func errResponse(rw http.ResponseWriter) {
	rw.WriteHeader(500)
	rw.Write([]byte("Error"))
}
