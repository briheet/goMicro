package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

var prices = map[string]float64{
	"ETH": 999.99,
	"BTC": 20000.0,
	"SY":  1000000.0,
}

type PriceService interface {
	FetchPrice(context.Context, string) (float64, error)
}

type priceService struct{}

func (s *priceService) FetchPrice(c context.Context, ticker string) (float64, error) {
	price, ok := prices[ticker]

	if !ok {
		return 0.0, fmt.Errorf("price for the ticker (%s) is not available", ticker)
	}

	return price, nil
}

type loggingService struct {
	next PriceService
}

func (s loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		reqID := ctx.Value("requestID")

		logrus.WithFields(logrus.Fields{
			"requestID": reqID,
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
			"ticker":    ticker,
		}).Info("FetchPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
