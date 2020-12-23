# Bitfinex API Client

A partial implementation of client for the Bitfinex API.

## Client

See [client.go](client.go).

## Binaries

Example

```
$ go run cmd/bitfinex-tickers/main.go -symbols tBSVUSD,tBTCUSD
{Symbol:tBSVUSD Bid:168.47 BidSize:1827.1089 Ask:168.77 AskSize:1131.6995 DailyChange:-1.61 DailyChangeRelative:-0.0095 LastPrice:168.69 Volume:2730.6577 High:173.5 Low:164.32}
{Symbol:tBTCUSD Bid:23587 BidSize:19.252996 Ask:23592 AskSize:70.17683 DailyChange:654 DailyChangeRelative:0.0285 LastPrice:23588 Volume:7314.4287 High:24013.668 Low:22348}
```

## References

- https://docs.bitfinex.com/docs


## License

The MIT License (MIT)

Copyright (c) 2020 Scott Barr

See [LICENSE](LICENSE)
