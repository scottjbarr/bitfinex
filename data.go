package bitfinex

import (
	"strconv"
	"strings"
)

func stripStringData(s string) string {
	r := strings.NewReplacer("[", "", "\"", "")
	return r.Replace(s)
}

func parseFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, err
	}

	return float32(f), nil
}
