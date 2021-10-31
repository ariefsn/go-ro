package core

type M map[string]interface{}
type Account string
type BaseUrl string
type PlaceType string
type QueryParam struct {
	Field string
	Value interface{}
}
type Options struct {
	AccountType Account
	ApiKey      string
}
