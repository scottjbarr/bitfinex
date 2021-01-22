package bitfinex

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"strings"
)

// Ticker represents a payload for a ticker for a trading pair.
type Ticker struct {
	Symbol              string  `json:"symbol,omitempty"`
	Bid                 float32 `json:"bid,omitempty"`
	BidSize             float32 `json:"bid_size,omitempty"`
	Ask                 float32 `json:"ask,omitempty"`
	AskSize             float32 `json:"ask_size,omitempty"`
	DailyChange         float32 `json:"daily_change,omitempty"`
	DailyChangeRelative float32 `json:"daily_change_rel,omitempty"`
	LastPrice           float32 `json:"last_price,omitempty"`
	Volume              float32 `json:"vol,omitempty"`
	High                float32 `json:"high,omitempty"`
	Low                 float32 `json:"low,omitempty"`
	Hash                string  `json:"-"`
}

// ParseTickers parses an array of Tickers.
func ParseTickers(body []byte) ([]Ticker, error) {
	// remove first/last square braces
	body = body[1 : len(body)-2]

	// split into records
	parts := bytes.Split(body, []byte("]"))

	records := [][]byte{}

	for i, r := range parts {
		if i == 0 {
			// first record is clean
			records = append(records, r)
			continue
		}

		// empty
		if len(r) == 0 {
			continue
		}

		// skip the first byte
		records = append(records, r[1:len(r)])
	}

	tickers := make([]Ticker, len(records), len(records))

	for i, r := range records {
		t, err := ParseTicker(r)
		if err != nil {
			return nil, err
		}

		tickers[i] = *t
	}

	return tickers, nil
}

// ParseTicker parses a single Ticker.
func ParseTicker(b []byte) (*Ticker, error) {
	parts := strings.Split(string(b), ",")

	bid, err := parseFloat32(parts[1])
	if err != nil {
		return nil, err
	}

	bidSize, err := parseFloat32(parts[2])
	if err != nil {
		return nil, err
	}

	ask, err := parseFloat32(parts[3])
	if err != nil {
		return nil, err
	}

	askSize, err := parseFloat32(parts[4])
	if err != nil {
		return nil, err
	}

	dailyChange, err := parseFloat32(parts[5])
	if err != nil {
		return nil, err
	}

	dailyChangeRelative, err := parseFloat32(parts[6])
	if err != nil {
		return nil, err
	}

	lastPrice, err := parseFloat32(parts[7])
	if err != nil {
		return nil, err
	}

	volume, err := parseFloat32(parts[8])
	if err != nil {
		return nil, err
	}

	high, err := parseFloat32(parts[9])
	if err != nil {
		return nil, err
	}

	low, err := parseFloat32(parts[10])
	if err != nil {
		return nil, err
	}

	// add a hash of the data to make it easy to check if a Ticker hash changed.
	hash := md5.Sum(b)

	t := Ticker{
		Symbol:              stripStringData(parts[0]),
		Bid:                 bid,
		BidSize:             bidSize,
		Ask:                 ask,
		AskSize:             askSize,
		DailyChange:         dailyChange,
		DailyChangeRelative: dailyChangeRelative,
		LastPrice:           lastPrice,
		Volume:              volume,
		High:                high,
		Low:                 low,
		Hash:                fmt.Sprintf("%x", hash),
	}

	return &t, nil
}
