package main

import (
	"context"
	"github.com/hashicorp/go-hclog"
	protos "github.com/idawud/go-gRpc-microservice/protos/currency"
	"google.golang.org/grpc"
)

func main()  {
	log := hclog.Default()
	// create client
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cc := protos.NewCurrencyClient(conn)
	rate, err := cc.GetRate(context.Background(), &protos.RateRequest{
		Base:        protos.Currencies_EUR,
		Destination: protos.Currencies_USD,
	})
	if err != nil {
		log.Error("error ", err)
		return
	}

	log.Info("Rate", "Rate", rate.Rate )
	log.Info("500 AT Rate", "info", rate.Rate * 500)
}