package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/briheet/micro/types"
)

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}

func MakeAPIFunc(fn APIFunc) http.HandlerFunc {
	ctx := context.Background()

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(ctx, "requestID", rand.Intn(10000))

		if err := fn(ctx, w, r); err != nil {
			writeJson(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

type JSONAPIServer struct {
	listenAddr string
	svc        PriceService
}

func NewJSONAPIServer(listenAddr string, svc PriceService) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", MakeAPIFunc(s.HandleFetchPrice))
	http.HandleFunc("/v2", MakeAPIFunc(s.HandleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func (s *JSONAPIServer) HandleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")
	if len(ticker) == 0 {
		return fmt.Errorf("invalid ticker")
	}

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	resp := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJson(w, http.StatusOK, resp)
}
