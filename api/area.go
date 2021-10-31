package api

import (
	"github.com/ariefsn/go-ro/core"
	"github.com/ariefsn/go-ro/models"
)

type ApiArea struct {
	opt core.Options
}

func NewApiArea(opt core.Options) *ApiArea {
	a := new(ApiArea)
	a.opt = opt

	return a
}

func (a *ApiArea) GetProvinces(id string) ([]models.Province, *core.Status, error) {
	datas := []models.Province{}
	status := &core.Status{}

	queryParams := core.GetQueryParams(core.QueryParam{
		Field: "id",
		Value: id,
	})

	uri := "/province" + queryParams

	result := core.Response{}
	resp, err := core.Client(a.opt, &result).Get(uri)

	status.Code = resp.StatusCode()
	status.Description = resp.Status()

	if err != nil {
		status.Description = err.Error()
		return datas, status, err
	}

	for _, p := range result.Results() {
		temp := models.Province{}
		core.MapToStruct(p.(map[string]interface{}), &temp)
		datas = append(datas, temp)
	}

	_status := result.Status()
	status = &_status

	return datas, status, nil
}

func (a *ApiArea) GetCities(province, id string) ([]models.City, *core.Status, error) {
	datas := []models.City{}
	status := &core.Status{}

	queryParams := core.GetQueryParams(core.QueryParam{
		Field: "id",
		Value: id,
	}, core.QueryParam{
		Field: "province",
		Value: province,
	})

	uri := "/city" + queryParams

	result := core.Response{}
	resp, err := core.Client(a.opt, &result).Get(uri)

	status.Code = resp.StatusCode()
	status.Description = resp.Status()

	if err != nil {
		status.Description = err.Error()
		return datas, status, err
	}

	for _, p := range result.Results() {
		temp := models.City{}
		core.MapToStruct(p.(map[string]interface{}), &temp)
		datas = append(datas, temp)
	}

	_status := result.Status()
	status = &_status

	return datas, status, nil
}

func (a *ApiArea) GetSubDistrict(city, id string) ([]models.SubDistrict, *core.Status, error) {
	datas := []models.SubDistrict{}
	status := &core.Status{}

	queryParams := core.GetQueryParams(core.QueryParam{
		Field: "id",
		Value: id,
	}, core.QueryParam{
		Field: "city",
		Value: city,
	})

	uri := "/subdistrict" + queryParams

	result := core.Response{}
	resp, err := core.Client(a.opt, &result).Get(uri)

	status.Code = resp.StatusCode()
	status.Description = resp.Status()

	if err != nil {
		status.Description = err.Error()
		return datas, status, err
	}

	for _, p := range result.Results() {
		temp := models.SubDistrict{}
		core.MapToStruct(p.(map[string]interface{}), &temp)
		datas = append(datas, temp)
	}

	_status := result.Status()
	status = &_status

	return datas, status, nil
}
