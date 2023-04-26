package handlers

import "net/http"

func HomePage(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello World!"))
}
