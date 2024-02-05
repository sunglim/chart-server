package chartserver

import (
	"fmt"
	"net/http"

	"github.com/cinar/indicator"
	"github.com/sunglim/chart-server/internal/database"
)

// There might be a better pattern for http routing.

func serverRun() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Rsi(w)
	})
	http.HandleFunc("/rsi", func(w http.ResponseWriter, r *http.Request) {
		Rsi(w)
	})

	fmt.Println(":8080")
	return http.ListenAndServe(":8080", nil)
}

func Rsi(w http.ResponseWriter) {
	db := Default()
	closing := database.GetClosingPrices(db.db)
	_, rsi := indicator.Rsi(closing)
	fmt.Fprintf(w, "Welcome! RSI: %f\n", len(rsi))
}
