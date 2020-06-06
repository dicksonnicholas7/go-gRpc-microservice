package server

import (
	"context"
	"github.com/hashicorp/go-hclog"
	data "github.com/idawud/go-gRpc-microservice/data"
	protos "github.com/idawud/go-gRpc-microservice/protos/currency"
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
