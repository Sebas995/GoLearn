package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	Id      int    "json:Id"
	Name    string "json:Name"
	Content string "json:Content"
}

type allTasks []task

var tasks = allTasks{
	{
		Id:      1,
		Name:    "Task One",
		Content: "Something",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ahre")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":8080", router))
}
