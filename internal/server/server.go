package server

import (
	"fmt"
	"net/http"
)

func Run() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})

	return http.ListenAndServe(":8080", nil)
}
