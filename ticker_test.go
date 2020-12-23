package bitfinex

import (
	"reflect"
	"testing"
)

func TestParseTickers(t *testing.T) {
	body := LoadFixture("tickers-multiple.txt")

	got, err := ParseTickers(body)
	if err != nil {
		t.Fatal(err)
	}

	want := []Ticker{
		{
			Symbol:              "tBTCUSD",
			Bid:                 23639,
			BidSize:             17.67616704,
			Ask:                 23640,
			AskSize:             31.781330249999993,
			DailyChange:         741,
			DailyChangeRelative: 0.0323,
			LastPrice:           23653,
			Volume:              7347.30314117,
			High:                24013.667964,
			Low:                 22348,
		},
		{
			Symbol:              "tBSVUSD",
			Bid:                 168.88,
			BidSize:             1648.9035091499998,
			Ask:                 168.96,
			AskSize:             1158.35702644,
			DailyChange:         -0.51,
			DailyChangeRelative: -0.003,
			LastPrice:           168.97,
			Volume:              2813.31065912,
			High:                173.5,
			Low:                 164.32,
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got\n%+v\nwant\n%+v", got, want)
	}
}
