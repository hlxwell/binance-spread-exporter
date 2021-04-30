package main

import (
	"fmt"
	"math"
	"time"
)

const FetchDuration = 1 * time.Second

// latest spread of each symbol
var LatestSymbolSpread = make(map[string]float64)

// all pairs spread delta.
var SpreadDelta = make(map[string]float64)

// Fetch all pairs
func DepthAutoFetcher(pairs []string) {
	client := BinanceClient{}

	for x := range time.Tick(FetchDuration) {
		fmt.Println("Updated at ", x)

		// actual fetching data from pair depth, and calculate spread and delta
		for _, pair := range pairs {
			depth, _ := client.FetchPairDepth(pair)
			SpreadDelta[pair] = math.Abs(LatestSymbolSpread[pair] - depth.Spread())
			LatestSymbolSpread[pair] = depth.Spread()
			fmt.Printf(
				"%s spread is %.8f delta %.8f. ",
				pair, depth.Spread(), SpreadDelta[pair],
			)
			fmt.Printf(
				"Top 200 Asks notional value: %.8f, Top 200 Bids notional value: %.8f \n",
				depth.TopNotionalValue(200, Asks),
				depth.TopNotionalValue(200, Bids),
			)
		}
	}
}
