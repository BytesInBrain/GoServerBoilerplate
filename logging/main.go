package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func loggingHandler(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Response) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL, Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index handler")
	fmt.Fprintf(w, "Welcome!")
}
func about(w http.ResponseWriter, r *http.Request) {
	log.Println("Execuiting index handler")
	fmt.Fprintf(w, "Welcome")
}
