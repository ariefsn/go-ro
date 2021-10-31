package api

import (
	"errors"

	"github.com/ariefsn/go-ro/core"
	"github.com/ariefsn/go-ro/models"
)

type ApiAwb struct {
	opt core.Options
}

func NewApiAwb(opt core.Options) *ApiAwb {
	a := new(ApiAwb)
	a.opt = opt

	return a
}

func (a *ApiAwb) GetAwb(awb, courier string) (*models.Awb, *core.Status, error) {
	status := &core.Status{}

	uri := "/waybill"

	result := core.Response{}
	resp, err := core.Client(a.opt, &result).SetFormData(map[string]string{
		"waybill": awb,
		"courier": courier,
	}).Post(uri)

	status.Code = resp.StatusCode()
	status.Description = resp.Status()

	if err != nil {
		status.Description = err.Error()
		return nil, status, err
	}

	_status := result.Status()
	status = &_status

	_data := result.Result()
	if _data == nil {
		return nil, status, errors.New("awb not found")
	}

	data := models.Awb{}

	err = core.MapToStruct(_data.(map[string]interface{}), &data)

	if err != nil {
		return nil, status, err
	}

	return &data, status, nil
}

func (a *ApiAwb) GetCouriers() []models.Courier {
	return []models.Courier{
		{
			Code:     "pos",
			Name:     "Pos Indonesia",
			Homepage: "https://www.posindonesia.co.id/",
		},
		{
			Code:     "jne",
			Name:     "JNE Express",
			Homepage: "https://www.jne.co.id/",
		},
		{
			Code:     "wahana",
			Name:     "Wahana Express",
			Homepage: "https://www.wahana.com/",
		},
		{
			Code:     "jnt",
			Name:     "Jnt Express",
			Homepage: "https://jet.co.id/",
		},
		{
			Code:     "sicepat",
			Name:     "SiCepat Express",
			Homepage: "https://www.sicepat.com/",
		},
		{
			Code:     "jet",
			Name:     "Jet Express",
			Homepage: "http://www.jetexpress.co.id/",
		},
		{
			Code:     "dse",
			Name:     "DSE",
			Homepage: "",
		},
		{
			Code:     "first",
			Name:     "First",
			Homepage: "",
		},
		{
			Code:     "ninja",
			Name:     "Ninja Xpress",
			Homepage: "https://www.ninjaxpress.co/",
		},
		{
			Code:     "lion",
			Name:     "Lion Parcel",
			Homepage: "https://lionparcel.com/",
		},
		{
			Code:     "idl",
			Name:     "IDL Cargo",
			Homepage: "http://www.idlcargo.co.id/web.idl/",
		},
		{
			Code:     "rex",
			Name:     "REX",
			Homepage: "https://www.rex.co.id/",
		},
		{
			Code:     "ide",
			Name:     "ID Express",
			Homepage: "https://idexpress.com/",
		},
		{
			Code:     "central",
			Name:     "Central Transport",
			Homepage: "https://www.centraltransport.com/",
		},
		{
			Code:     "anteraja",
			Name:     "Anteraja",
			Homepage: "https://anteraja.id/",
		},
	}
}
