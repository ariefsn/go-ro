package models

type City struct {
	Id         string `json:"city_id"`
	Name       string `json:"city_name"`
	Province   string `json:"province"`
	ProvinceId string `json:"province_id"`
	Type       string `json:"type"`
	PostalCode string `json:"postal_code"`
}
