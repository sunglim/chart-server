package server

import (
	"fmt"
	"net/http"
)

func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})

	http.ListenAndServe(":8080", nil)
}
