package main

import (
	"fmt"
	"log"
	"net/http"
	"recipes/data"
	"recipes/handlers"
)

func main() {

	data.FetchAllRecipes()
	http.HandleFunc("/", handlers.HomePage)
	fmt.Println(data.AllRecipes)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
