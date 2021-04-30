package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type Ticker struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int    `json:"openTime"`
	CloseTime          int    `json:"closeTime"`
	FirstId            int    `json:"firstId"`
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

type TickerList struct {
	Tickers []Ticker
}

type SymbolVol struct {
	Symbol      string
	QuoteVolume float64
}

// List Top 5 volume in given symbol.
// Usage: Top5VolIn("BTC")
func (list TickerList) Top5VolIn(baseSymbol string) []*SymbolVol {
	symbolVolHash := make(map[string]*SymbolVol)
	var vol float64
	var isRightBase bool

	for _, ticker := range list.Tickers {
		// Filter tickers priced by Base Symbol
		isRightBase, _ = regexp.MatchString(fmt.Sprintf("%s$", baseSymbol), ticker.Symbol)
		if !isRightBase {
			continue
		}

		vol, _ = strconv.ParseFloat(ticker.QuoteVolume, 64)
		if symbolVolHash[ticker.Symbol] == nil {
			symbolVolHash[ticker.Symbol] = &SymbolVol{ticker.Symbol, vol}
			continue
		}

		symbolVolHash[ticker.Symbol].QuoteVolume += vol
	}

	sortedSymbolVols := sortedHashByValue(symbolVolHash)
	return sortedSymbolVols[:5] // only took top 5 elements.
}

// Sort hash value and return a sorted array.
func sortedHashByValue(hash map[string]*SymbolVol) []*SymbolVol {
	var symbolVols []*SymbolVol
	// push all values into array
	for _, vol := range hash {
		symbolVols = append(symbolVols, vol)
	}

	// sort the array by volume.
	sort.SliceStable(symbolVols, func(i, j int) bool {
		return symbolVols[i].QuoteVolume > symbolVols[j].QuoteVolume
	})

	return symbolVols
}
