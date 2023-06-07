package main

import (
	"net/http"
)

func bootstrap() *http.ServeMux {
	// 1. setup mux (HTTP request router)
	// 2. setup hadler (handler for request)
	// 3. connect mux and handler
	mux := http.NewServeMux()
	filehandler := http.FileServer(http.Dir("public"))
	mux.Handle("/", filehandler)
	return mux
}

// curl html://localhost:8080/welcome.html
func main() {
	http.ListenAndServe(":8080", bootstrap())
}
