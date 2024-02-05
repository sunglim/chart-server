package chartserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cinar/indicator"
	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/pkg/cmd/chartserver"
	"gorm.io/gorm"
)

// There might be a better pattern for http routing.
func serverRun(opts *chartserver.Options) error {
	db := CreateDatabase(opts.BaseDirectory + "/database.db")

	http.HandleFunc("/rsi", func(w http.ResponseWriter, r *http.Request) {
		handleRsi(db.db, w, r)
	})

	addr := ":" + strconv.Itoa(opts.Port)
	fmt.Printf("Serving on %s...", addr)
	return http.ListenAndServe(addr, nil)
}

func handleRsi(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	if symbol == "" {
		fmt.Fprintf(w, "No symbol provided.")
		return
	}

	closing := database.GetClosingPrices(db, symbol)
	_, rsi := indicator.Rsi(closing)
	fmt.Fprintf(w, "Welcome! RSI: %f\n", rsi[len(rsi)-1])
}
