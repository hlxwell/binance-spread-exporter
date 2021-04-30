# README

## Diagram

![](https://raw.githubusercontent.com/hlxwell/binance-spread-exporter/master/docs/diagram.png)

## How to start

```
make run
```

## How bulid

```
make build
```


## PromQL

Metrics provided at:

http://localhost:8080/metrics

Output format of exporter

```
spread{symbol="BTCUSDT"} 0.01000000
delta{symbol="BTCUSDT"} 0.00000000
spread{symbol="BNBUSDT"} 0.01000000
delta{symbol="BNBUSDT"} 0.00000000
spread{symbol="ETHUSDT"} 0.01000000
delta{symbol="ETHUSDT"} 0.00000000
spread{symbol="BUSDUSDT"} 0.00010000
delta{symbol="BUSDUSDT"} 0.00000000
spread{symbol="DOGEUSDT"} 0.00001000
delta{symbol="DOGEUSDT"} 0.00000000
```

- List all the spread of a specific pair.
```
spread{symbol="BTCUSDT"}
```

- List all the delta of a specific pair.
```
delta{symbol="BTCUSDT"}
```

## TODO

1. API auto failure over.

## API is used

1. https://api.binance.com/api/v3/ticker/24hr
2. https://api.binance.com/api/v3/depth?symbol=LTCBTC
