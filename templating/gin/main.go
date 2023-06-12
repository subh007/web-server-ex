package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID       int
	Name     string
	Status   string
	Priority int
}

var (
	tasks = []Task{
		{
			ID:       1,
			Name:     "Bob",
			Status:   "Not Started",
			Priority: 5,
		}, {
			ID:       2,
			Name:     "Steven",
			Status:   "Done",
			Priority: 3,
		}, {
			ID:       3,
			Name:     "Rob",
			Status:   "Review",
			Priority: 4,
		},
	}
)

func indexPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "task.tmpl", gin.H{
		"Tasks": tasks,
	})
}

func main() {
	router := gin.Default()
	router.GET("/viewtasks", indexPageHandler)
	router.LoadHTMLGlob("template/*")
	router.Run(":8080")
}
