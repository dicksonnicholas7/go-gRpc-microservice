package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/idawud/go-gRpc-microservice/data"
	protos "github.com/idawud/go-gRpc-microservice/protos/currency"
	"github.com/idawud/go-gRpc-microservice/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRate(log)
	if err != nil {
		log.Error("Unable to generate rates", "error", err)
		os.Exit(1)
	}
	gs := grpc.NewServer()

	cs := server.NewCurrency(log, rates)
	protos.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	// run server
	listen, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen ", " error", err)
		os.Exit(1)
	}
	gs.Serve(listen)
}
