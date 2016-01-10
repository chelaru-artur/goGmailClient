package main

import (
	"log"
	"net/http"
)

func main() {

	//Gmail()

	router := NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	log.Println("Server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
