package main

import (
	"fmt"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("recived request for about")
	fmt.Fprintf(w, "This is go-web framework session")
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("==> url: [%s]\n", r.URL)
		next.ServeHTTP(w, r)
		fmt.Printf("=== finish ===\n")
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/about", LoggerMiddleware(http.HandlerFunc(AboutHandler)))
	http.ListenAndServe(":8080", mux)
}
