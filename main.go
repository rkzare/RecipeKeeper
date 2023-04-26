package main

import (
	"log"
	"net/http"
	"recipes/handlers"
)

func main() {

	http.HandleFunc("/", handlers.HomePage)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
