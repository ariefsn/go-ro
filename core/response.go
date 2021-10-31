package core

import (
	"reflect"

	"github.com/ariefsn/go-ro/models"
)

type Response struct {
	RajaOngkir struct {
		Query       interface{} `json:"query"`
		Status      Status      `json:"status"`
		Result      interface{} `json:"result"`
		Results     interface{} `json:"results"`
		Origin      models.City `json:"origin_details"`
		Destination models.City `json:"destination_details"`
	} `json:"rajaongkir"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

func (r *Response) Origin() models.City {
	return r.RajaOngkir.Origin
}

func (r *Response) Destination() models.City {
	return r.RajaOngkir.Destination
}

func (r *Response) Status() Status {
	return r.RajaOngkir.Status
}

func (r *Response) Query() map[string]interface{} {
	query := r.RajaOngkir.Query

	if query == nil {
		return nil
	}

	dataType := reflect.TypeOf(query).Kind()

	if dataType == reflect.Map {
		return query.(map[string]interface{})
	}

	return nil
}

func (r *Response) Result() interface{} {
	res := r.RajaOngkir.Result

	if res == nil {
		return nil
	}

	return res
}

func (r *Response) Results() []interface{} {
	q := []interface{}{}

	res := r.RajaOngkir.Results

	if res == nil {
		return q
	}

	dataType := reflect.TypeOf(res).Kind()

	switch dataType {
	case reflect.Map:
		q = append(q, res)
	case reflect.Slice:
		q = res.([]interface{})
	}

	return q
}
