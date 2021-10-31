package core

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

func Client(opt Options, result *Response) *resty.Request {
	baseUrl := GetBaseUrl(opt.AccountType)

	c := resty.New()
	c.SetHeader("key", opt.ApiKey)
	c.SetHostURL(string(baseUrl))

	if result != nil {
		return c.R().EnableTrace().SetResult(result)
	}

	return c.R().EnableTrace()
}

func GetBaseUrl(accountType Account) BaseUrl {
	switch accountType {
	case Starter:
		return StarterUrl
	case Basic:
		return BasicUrl
	case Pro:
		return ProUrl
	default:
		return StarterUrl
	}
}

func GetQueryParams(params ...QueryParam) string {
	if len(params) == 0 {
		return ""
	}

	queryParams := []string{}

	for _, p := range params {
		if p.Value != nil {
			param := fmt.Sprintf("%s=%s", p.Field, p.Value)
			queryParams = append(queryParams, param)
		}
	}

	return "?" + strings.Join(queryParams, "&")
}

func MapToStruct(m map[string]interface{}, target interface{}) error {
	data, err := json.Marshal(m)

	if err != nil {
		return err
	}

	return json.Unmarshal(data, target)
}

func ByteToStruct(b []byte, target interface{}) error {
	return json.Unmarshal(b, target)
}
