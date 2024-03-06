package chartserver

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cinar/indicator"
	"github.com/sunglim/chart-server/internal/database"
	"github.com/sunglim/chart-server/pkg/cmd/chartserver"
	"gorm.io/gorm"
	"time"
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

	date := r.URL.Query().Get("date")
	if date == "" || date == "today" {
		date = time.Now().Format("2006-01-02")
	}

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Fprintf(w, "Invalid date provided.")
		return
	}

	rsi := getRsiFromDatabase(date, symbol, db)
	fmt.Fprintf(w, "Welcome! %s, RSI: %f\n", date, rsi)
}

func getRsiFromDatabase(date, symbol string, db *gorm.DB) float64 {
	closing := database.GetClosingPricesOnDate(db, symbol, date)
	_, rsi := indicator.Rsi(closing)
	if len(rsi) == 0 {
		return 0
	}

	return rsi[len(rsi)-1]
}