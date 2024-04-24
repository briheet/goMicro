package proto

import (
	"context"

	"github.com/briheet/micro/types"
	"google.golang.org/grpc"
)

type PriceFetcherClient interface {
	FetchPrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*types.PriceResponse, error)
}

type priceFethcerCLient struct {
	cc grpc.ClientConnInterface
}

func NewPriceFetcherClient(cc grpc.ClientConnInterface) PriceFetcherClient {
	return &priceFethcerCLient{cc}
}

func (c *priceFethcerCLient) FetchPrice(ctx context.Context, in *PriceRequest, opts ...grpc.CallOption) (*types.PriceResponse, error) {
	out := new(PriceResponse)
	return nil, nil
}
