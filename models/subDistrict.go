package models

type SubDistrict struct {
	Id           string `json:"subdistrict_id"`
	Name         string `json:"subdistrict_name"`
	ProvinceName string `json:"province"`
	ProvinceId   string `json:"province_id"`
	CityName     string `json:"city"`
	CityId       string `json:"city_id"`
	Type         string `json:"type"`
}
