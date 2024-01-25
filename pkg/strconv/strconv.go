package strconv

import "strconv"

func ParseFloat64OrZero(s string) float64 {
	float64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}

	return float64

}

func ParseInt64OrZero(s string) int64 {
	int64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}

	return int64
}
