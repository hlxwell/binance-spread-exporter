package main

import (
	"fmt"
)

// latest spread of each symbol
var LatestSymbolSpread = make(map[string]float64)

// all pairs spread delta.
var SpreadDelta = make(map[string]float64)

func main() {
	// Top 5 volume symbol.
	var top5BtcPairs []string
	var top5UsdtPairs []string

	client := BinanceClient{}
	tickerList, _ := client.FetchDayTickerList()

	// Print TOP 5 BTC pairs
	fmt.Println("TOP 5 BTC pairs:")
	for _, vol := range tickerList.Top5VolIn("BTC") {
		top5BtcPairs = append(top5BtcPairs, vol.Symbol)
		fmt.Printf("- %s: %.8f\n", vol.Symbol, vol.QuoteVolume)
	}

	// Print TOP 5 USDT pairs
	fmt.Println("TOP 5 USDT pairs:")
	for _, vol := range tickerList.Top5VolIn("USDT") {
		top5UsdtPairs = append(top5UsdtPairs, vol.Symbol)
		fmt.Printf("- %s: %.8f\n", vol.Symbol, vol.QuoteVolume)
	}

	// Print Spread of USDT pairs
	go DepthAutoFetcher(top5UsdtPairs)

	// Start the metrics server.
	StartPromServer()
}
