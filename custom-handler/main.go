package main

import (
	"fmt"
	"net/http"
)

var (
	Participants []string = []string{
		"Ram",
		"Mohan",
	}
	Tasks []string = []string{
		"complete assigment",
		"submit assigment",
	}
)

type TaskHandler struct {
}

func (t *TaskHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, task := range Tasks {
		fmt.Fprintf(w, task+"\n")
	}
}

func ParticipantList(w http.ResponseWriter, r *http.Request) {
	for _, participant := range Participants {
		fmt.Fprintf(w, participant+"\n")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/task", &TaskHandler{})
	mux.HandleFunc("/participant", ParticipantList)
	http.ListenAndServe(":8080", mux)
}
