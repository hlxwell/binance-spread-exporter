package main

import "testing"

func TestFetchDayTickerList(t *testing.T) {
	client := &BinanceClient{}
	tickers, err := client.FetchDayTickerList()

	if err != nil || len(tickers) == 0 {
		t.Error("binance client cannot fetch 24 hrs tickers")
	}
}

func TestFetchPairDepth(t *testing.T) {
	client := &BinanceClient{}
	depth, err := client.FetchPairDepth("BTCUSDT")

	if err != nil || len(depth.Asks) == 0 || len(depth.Bids) == 0 {
		t.Error("binance client cannot fetch pair depth")
	}
}
