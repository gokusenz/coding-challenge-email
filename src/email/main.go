package main

import (
	"net/http"
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
	http.HandleFunc("/email", responseEmail)
	http.ListenAndServe(":80", nil)
}

func errResponse(rw http.ResponseWriter) {
	rw.WriteHeader(500)
	rw.Write([]byte("Error"))
}
