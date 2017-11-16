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
	"github.com/justinas/alice"
)

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func loggingHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Welcome!")
}

func emailHandler(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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

	commonHandlers := alice.New(loggingHandler, recoverHandler)
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/email", commonHandlers.ThenFunc(emailHandler))
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
