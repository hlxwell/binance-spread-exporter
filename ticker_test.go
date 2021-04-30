package main

import "testing"

func TestTickerListLoad(t *testing.T) {
	tickerList := &TickerList{}
	tickerList.Load()

	if len(tickerList.Tickers) <= 0 {
		t.Error("TickerList.Load cannot load data.")
	}
}

func TestTop5VolIn(t *testing.T) {
	tickerList := &TickerList{}
	tickerList.Load()
	result := tickerList.Top5VolIn("BTC")
	if len(result) <= 0 {
		t.Error("Top5VolIn BTC cannot get data.")
	}
}
