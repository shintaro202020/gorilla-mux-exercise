package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func HelloHandler(res http.ResponseWriter, req *http.Request) {}
