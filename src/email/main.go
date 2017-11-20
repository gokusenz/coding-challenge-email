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

// emailHandler
func emailHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		defer r.Body.Close()
		var e *mail.Email
		if r.Body == nil {
			http.Error(rw, "Please send a request body", 400)
			return
		}
		err := json.NewDecoder(r.Body).Decode(&e)
		if err != nil {
			http.Error(rw, err.Error(), 400)
			return
		}
		el := mail.EmailInfoer{}
		resp, err := mail.Send(el, e)
		if err != nil {
			log.Println(err)
		}
		log.Println(resp)
		rw.WriteHeader(resp)
		rw.Write([]byte("Success"))

	} else {
		http.Error(rw, "Invalid request method", http.StatusMethodNotAllowed)
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

}

func errResponse(rw http.ResponseWriter) {
	rw.WriteHeader(500)
	rw.Write([]byte("Error"))
}
