# README

## Diagram

![](https://raw.githubusercontent.com/hlxwell/binance-spread-exporter/master/docs/diagram.png)

## How to start

```
make run
```

or

```
make build 
```

## API is used

1. https://api.binance.com/api/v3/ticker/24hr
2. https://api.binance.com/api/v3/depth?symbol=LTCBTC

## PromQL

- List all the spread of a specific pair.
```
spread{symbol="BTCUSDT"}
```

- List all the delta of a specific pair.
```
delta{symbol="BTCUSDT"}
```
