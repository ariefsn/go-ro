package api

import (
	"errors"

	"github.com/ariefsn/go-ro/core"
	"github.com/ariefsn/go-ro/models"
)

type ApiCurrency struct {
	opt core.Options
}

func NewApiCurrency(opt core.Options) *ApiCurrency {
	a := new(ApiCurrency)
	a.opt = opt

	return a
}

func (a *ApiCurrency) GetCurrency() (*models.Currency, *core.Status, error) {
	status := &core.Status{}

	uri := "/currency"

	result := core.Response{}
	resp, err := core.Client(a.opt, &result).Get(uri)

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
		return nil, status, errors.New("currency not found")
	}

	data := models.Currency{}

	err = core.MapToStruct(_data.(map[string]interface{}), &data)

	if err != nil {
		return nil, status, err
	}

	return &data, status, nil
}
