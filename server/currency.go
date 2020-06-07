package server

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/idawud/go-gRpc-microservice/data"
	protos "github.com/idawud/go-gRpc-microservice/protos/currency"
	"io"
	"time"
)

type Currency struct {
	log hclog.Logger
	rates *data.ExchangeRates
}

func NewCurrency(l hclog.Logger, r *data.ExchangeRates)  *Currency{
	return &Currency{log: l, rates: r}
}
func (c *Currency) GetRate( ctx context.Context, rr *protos.RateRequest) (*protos.RateResponse, error) {
	c.log.Info("Handle GetRate", "base", rr.GetBase(), "destination", rr.GetDestination())
	rate, err := c.rates.GetRate(rr.GetBase().String(), rr.GetDestination().String())
	if err != nil {
		return nil, err
	}
	return &protos.RateResponse{Rate: rate}, nil
}

func (c *Currency) SubscribeRates(src protos.Currency_SubscribeRatesServer) error {
	// listen for receive stream to process
	go func() {
		for {
			recv, err := src.Recv()
			if err == io.EOF {
				c.log.Info("Client connection closed")
				break
			}

			if err != nil {
				c.log.Error("Unable to read from client", "error", err)
				break
			}
			req := fmt.Sprintf("{ Base: %s, Destination: %s}", recv.GetBase().String(), recv.GetDestination().String())
			c.log.Info("Handle client request", "request", req)
		}
	}()

	// server stream every 5sec
	for  {
		rate, err := c.rates.GetRate(protos.Currencies_EUR.String(), protos.Currencies_USD.String())
		if err != nil {
			return err
		}
		err = src.Send(&protos.RateResponse{Rate: rate })
		if err != nil {
			return err
		}
		time.Sleep(5 * time.Second)
	}
}