package main

import (
	"net"

	"github.com/briheet/micro/proto"
	"google.golang.org/grpc"
)

func makeGRPCServerandRun(listenAddr string, svc PriceService) error {
	grpcPriceFetcher := NewGRPCPriceFetcherServer(svc)

	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	return server.Serve(ln)
}
