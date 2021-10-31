package models

type CostValue struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type CostService struct {
	Service     string      `json:"service"`
	Description string      `json:"description"`
	Cost        []CostValue `json:"cost"`
}

type Cost struct {
	Code  string        `json:"code"`
	Name  string        `json:"name"`
	Costs []CostService `json:"costs"`
}
