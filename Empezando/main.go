package main

import (
	"log"
	"net/http"
	""

	"github.com/gorilla/mux"
)

var ahre modelo

func indexRoute(w http.ResponseWriter, r *http.Request) {
	var ahre modelo
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":8080", router))
}
