# GoRo

> Golang package for RajaOngkir API

## Install

```go
go get github.com/ariefsn/go-ro
```

## How to use

1. Create instance

    ```go
    package main

    import (
      "fmt"
      goro "github.com/ariefsn/go-ro"
      "github.com/ariefsn/go-ro/core"
    )

    ro := goro.New(core.Options{
      AccountType: "YOUR ACCOUNT TYPE",
      ApiKey:      "YOUR API KEY",
    })
    ```

2. Use it

     - Couriers

        ```go
        couriers := ro.Awb().GetCouriers()

        fmt.Println("couriers", couriers)
        ```

     - Currency

        ```go
        currency, status, err := ro.Currency().GetCurrency()

        fmt.Println("currency", currency, status, err)
        ```

     - Province

        ```go
        prov, status, err := ro.Area().GetProvinces("12")

        fmt.Println("prov", prov, status, err)
        ```

     - City

        ```go
        city, status, err := ro.Area().GetCities("5", "501")

        fmt.Println("city", city, status, err)
        ```

     - SubDistrict

        ```go
        subDistrict, status, err := ro.Area().GetSubDistrict("501", "6994")

        fmt.Println("subDistrict", subDistrict, status, err)
        ```

     - Costs

        ```go
        costParams := core.CostPayload{
          Origin:          "501",
          OriginType:      core.PlaceTypeCity,
          Destination:     "114",
          DestinationType: core.PlaceTypeSubDistrict,
          Weight:          1700,
          Couriers:        []string{"jne", "jnt"},
        }

        origin, destination, cost, status, err := ro.Cost().GetCosts(costParams)

        fmt.Println("cost", origin, destination, cost, status, err)
        ```

     - AWB

        ```go
        awb, status, err := ro.Awb().GetAwb("10002610261101", "anteraja")

        fmt.Println("awb", awb, status, err)
        ```
