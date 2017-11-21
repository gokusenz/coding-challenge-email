package main

import (
	"email/mail"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

// emailHandler
func emailHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Access-Control-Allow-Origin", "https://coding-challenge-email.firebaseapp.com")
	rw.Header().Set("Access-Control-Allow-Methods", "POST")
	if r.Method == "POST" {
		defer r.Body.Close()
		var e *mail.Email
		if r.Body == nil {
			log.Printf("Please send a request body")
			http.Error(rw, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			http.Error(rw, err.Error(), 400)
			return
		}
		el := mail.EmailInfoer{}
		respCode, respMsg := mail.Send(el, e)
		if err != nil {
			log.Println(err)
		}
		log.Printf("[CODE] %v : %v", respCode, respMsg)
		rw.Write([]byte(respMsg))

	} else {
		log.Printf("Invalid request method")
		http.Error(rw, "Invalid request method", http.StatusMethodNotAllowed)
	}

}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	commonHandlers := alice.New(loggingHandler, recoverHandler)
	http.Handle("/", commonHandlers.ThenFunc(indexHandler))
	http.Handle("/email", commonHandlers.ThenFunc(emailHandler))
	http.ListenAndServe(":8080", nil)

}
