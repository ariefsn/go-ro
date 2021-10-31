package goro

import (
	"github.com/ariefsn/go-ro/api"
	"github.com/ariefsn/go-ro/core"
)

type RajaOngkir struct {
	Options core.Options
}

func New(opt core.Options) *RajaOngkir {
	ro := new(RajaOngkir)
	ro.Options = opt

	return ro
}

func (r *RajaOngkir) Area() *api.ApiArea {
	return api.NewApiArea(r.Options)
}

func (r *RajaOngkir) Cost() *api.ApiCost {
	return api.NewApiCost(r.Options)
}

func (r *RajaOngkir) Currency() *api.ApiCurrency {
	return api.NewApiCurrency(r.Options)
}

func (r *RajaOngkir) Awb() *api.ApiAwb {
	return api.NewApiAwb(r.Options)
}
