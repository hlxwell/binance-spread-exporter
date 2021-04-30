package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// TODO api url failure over.

// Binance client
type BinanceClient struct{}

// Fetch ticker by 24 hrs.
func (client BinanceClient) FetchDayTickerList() ([]Ticker, error) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/24hr")
	if err != nil {
		fmt.Println("Error when fetching top vol: ", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when reading body: ", err)
		return nil, err
	}

	var tickers []Ticker
	json.Unmarshal(body, &tickers)
	return tickers, nil
}

// Fetch depth in order book by given symbol.
func (client BinanceClient) FetchPairDepth(symbol string) (Depth, error) {
	url := fmt.Sprintf("https://api.binance.com/api/v3/depth?symbol=%s&limit=500", symbol)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error when fetching pair depth: ", err)
		return Depth{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error when reading body: ", err)
		return Depth{}, err
	}

	depth := Depth{}
	json.Unmarshal(body, &depth)
	return depth, nil
}
