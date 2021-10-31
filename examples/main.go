package main

import (
	"fmt"

	goro "github.com/ariefsn/go-ro"
	"github.com/ariefsn/go-ro/core"
)

func main() {
	ro := goro.New(core.Options{
		AccountType: core.Pro,
		ApiKey:      "c8c6d68deeadcf379f25902b6bf7fbd0",
	})

	resCurrency, resStatus, err := ro.Currency().GetCurrency()

	fmt.Println("currency", resCurrency, err)

	resProv, resStatus, err := ro.Area().GetProvinces("12")

	fmt.Println("prov", resProv, err)

	resCity, resStatus, err := ro.Area().GetCities("5", "501")

	fmt.Println("city", resCity, err)

	resSubDistrict, resStatus, err := ro.Area().GetSubDistrict("501", "6994")

	fmt.Println("subDistrict", resSubDistrict, err)

	costParams := core.CostPayload{
		Origin:          "501",
		OriginType:      core.PlaceTypeCity,
		Destination:     "114",
		DestinationType: core.PlaceTypeSubDistrict,
		Weight:          1700,
		Couriers:        []string{"jne", "jnt"},
	}

	originCost, destinationCost, resCost, resStatus, err := ro.Cost().GetCosts(costParams)

	fmt.Println("cost", originCost, destinationCost, resCost, resStatus, err)

	resAwb, resStatus, err := ro.Awb().GetAwb("10002610261101", "anteraja")

	fmt.Println("awb", resAwb, resStatus, err)

	resCouriers := ro.Awb().GetCouriers()

	fmt.Println("couriers", resCouriers)
}
