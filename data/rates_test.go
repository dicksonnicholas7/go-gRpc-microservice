package data

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"testing"
)

func TestNewRates(t *testing.T) {
	tr, err := NewRate(hclog.Default())

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v", tr.rates)
}
