package main

import "context"

var prices = map[string]float64{
	"ETH": 999.99,
	"BTC": 20000.0,
	"SY":  1000000.0,
}

type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}
