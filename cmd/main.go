package main

import (
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	mux := gin.Default()

	return mux
}

// Request -> Route Parser -> [Optional Middleware] -> Route Handler -> [Optional Middleware] -> Response
func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
