package main

import (
	"fmt"
	"net/http"
)

func promDataHandler(w http.ResponseWriter, r *http.Request) {
	for pair, delta := range SpreadDelta {
		fmt.Fprintf(w, "spread{symbol=\"%s\"} %.8f \n", pair, LatestSymbolSpread[pair])
		fmt.Fprintf(w, "delta{symbol=\"%s\"} %.8f \n", pair, delta)
	}
}

func StartPromServer() {
	http.HandleFunc("/metrics", promDataHandler)
	http.ListenAndServe(":8080", nil)
}
