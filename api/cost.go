package api

import (
	"errors"
	"strconv"
	"strings"

	"github.com/ariefsn/go-ro/core"
	"github.com/ariefsn/go-ro/models"
)

type ApiCost struct {
	opt core.Options
}

func NewApiCost(opt core.Options) *ApiCost {
	a := new(ApiCost)
	a.opt = opt

	return a
}

func (a *ApiCost) GetCosts(params core.CostPayload) (origin *models.City, destination *models.City, costs []models.Cost, status *core.Status, err error) {
	uri := "/cost"

	result := core.Response{}
	payload := map[string]string{
		"origin":      params.Origin,
		"destination": params.Destination,
		"courier":     strings.Join(params.Couriers, ":"),
		"weight":      strconv.Itoa(params.Weight),
	}

	if a.opt.AccountType == core.Pro {
		if params.OriginType == "" || params.DestinationType == "" {
			return nil, nil, nil, nil, errors.New("pro account need to define origin and destination type")
		}

		payload["originType"] = string(params.OriginType)
		payload["destinationType"] = string(params.DestinationType)
	}

	resp, err := core.Client(a.opt, &result).SetFormData(payload).Post(uri)

	status = &core.Status{
		Code:        resp.StatusCode(),
		Description: resp.Status(),
	}

	if err != nil {
		status.Description = err.Error()
		return
	}

	// byteToStruct(resp.Body(), &result)

	for _, p := range result.Results() {
		temp := models.Cost{}
		core.MapToStruct(p.(map[string]interface{}), &temp)
		costs = append(costs, temp)
	}

	_origin := result.Origin()
	_destination := result.Destination()

	origin = &_origin
	destination = &_destination

	return
}
