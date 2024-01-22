package controller

// import "github.com/cinar/indicator"
import (
	"fmt"

	"github.com/cinar/indicator"
)

func HandleStockStatistics() {
	// declare float64 clice
	//closingPrices := []float64{398.69, 393.87, 389.47, 390.27, 388.47, 384.63, 382.77, 375.79, 374.69, 367.75, 367.94, 370.60, 370.87, 376.04, 375.28}

	closingPrices2 := []float64{375.28, 376.04, 370.87, 370.60, 367.94, 367.75, 374.69, 375.79, 382.77, 384.63, 388.47, 390.27, 389.47, 393.87, 398.69}

	rs, rsi := indicator.Rsi(closingPrices2)
	fmt.Print(rsi)
	fmt.Print(rs)
	// Read last 14 days closing price
}

func InsertHistoricalData() {

}
