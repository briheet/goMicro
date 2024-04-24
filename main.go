package main

import (
	"context"
	"flag"
	"log"

	"github.com/briheet/micro/client"
)

func main() {
	var (
		jsonAddr = flag.String("json", ":3000", "listen address of the json transport")
		grpcAddr = flag.String("grpc", ":4000", "listen address of the grpc transport")
		svc      = loggingService{&priceService{}}
		ctx      = context.Background()
	)
	flag.Parse()

	grpcClient, err := client.NewGRPCClient(":4000")
	if err != nil {
		log.Fatal(err)
	}
}
