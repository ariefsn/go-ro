package models

type Currency struct {
	Value      float64 `json:"value"`
	SourceName string  `json:"source_name"`
	SourceLink string  `json:"source_link"`
	UpdatedAt  string  `json:"last_update"`
}
