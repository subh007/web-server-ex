package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "welcome everyone",
	})
}

func main() {
	router := gin.Default()
	router.GET("/index:", indexPageHandler)
	router.LoadHTMLGlob("template/*")
	router.Run(":8080")
}
