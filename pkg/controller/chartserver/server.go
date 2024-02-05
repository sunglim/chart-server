package chartserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cinar/indicator"
	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/pkg/cmd/chartserver"
)

// There might be a better pattern for http routing.
func serverRun(opts *chartserver.Options) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Rsi(w)
	})
	http.HandleFunc("/rsi", func(w http.ResponseWriter, r *http.Request) {
		Rsi(w)
	})

	addr := ":" + strconv.Itoa(opts.Port)
	fmt.Printf("Serving on %s...", addr)
	return http.ListenAndServe(addr, nil)
}

func Rsi(w http.ResponseWriter) {
	db := Default()
	closing := database.GetClosingPrices(db.db)
	_, rsi := indicator.Rsi(closing)
	fmt.Fprintf(w, "Welcome! RSI: %f\n", len(rsi))
}
