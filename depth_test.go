package main

import (
	"testing"
)

func TestTopNotionalValue(t *testing.T) {
	depth := &Depth{}
	depth.Load("XRPUSDT")
	bidsNotionalVal := depth.TopNotionalValue(100, Bids)
	asksNotionalVal := depth.TopNotionalValue(100, Asks)

	if asksNotionalVal <= 0 || bidsNotionalVal <= 0 {
		t.Errorf("Cannot get Top Notional values, it was %.8f, %.8f", bidsNotionalVal, asksNotionalVal)
	}
}

func TestSpread(t *testing.T) {
	depth := &Depth{}
	depth.Load("XRPUSDT")
	spread := depth.Spread()
	if spread < 0 {
		t.Errorf("Something wrong with Spread() it was %.8f, which is less than 0", spread)
	}
}
