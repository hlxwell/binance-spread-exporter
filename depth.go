package main

import (
	"strconv"
)

type PriceVol []string

type Depth struct {
	Bids []PriceVol `json:"bids"`
	Asks []PriceVol `json:"asks"`
}

type OrderType int

const (
	Bids OrderType = 0
	Asks OrderType = 1
)

func (depth *Depth) Load(pair string) {
	client := BinanceClient{}
	newDepth, _ := client.FetchPairDepth(pair)
	depth.Asks = newDepth.Asks
	depth.Bids = newDepth.Bids
}

// TopNotionalValue
func (depth Depth) TopNotionalValue(limit int, orderType OrderType) float64 {
	if len(depth.Bids) < 200 || len(depth.Asks) < 200 {
		return 0.0
	}

	var sum float64 = 0
	var price, vol float64
	var orders []PriceVol

	if orderType == Bids {
		orders = depth.Bids[0:200]
	} else {
		orders = depth.Asks[0:200]
	}

	for _, priceVol := range orders {
		price, _ = strconv.ParseFloat(priceVol[0], 64)
		vol, _ = strconv.ParseFloat(priceVol[1], 64)
		sum += price * vol
	}
	return sum
}

func (depth Depth) Spread() float64 {
	if len(depth.Asks) == 0 {
		return 0.0
	}

	lowestAsk, _ := strconv.ParseFloat(depth.Asks[0][0], 64)
	highestBid, _ := strconv.ParseFloat(depth.Bids[0][0], 64)
	return lowestAsk - highestBid
}
