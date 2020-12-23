package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/scottjbarr/bitfinex"
)

func main() {
	s := ""

	flag.StringVar(&s, "symbols", "", "e.g. tBSVUSD,tBTCUSD")
	flag.Parse()

	symbols := strings.Split(s, ",")

	if len(symbols) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	c := bitfinex.New()

	tickers, err := c.Tickers(symbols)
	if err != nil {
		panic(err)
	}

	for _, t := range tickers {
		fmt.Printf("%+v\n", t)
	}
}
