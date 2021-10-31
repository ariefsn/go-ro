package core

type CostPayload struct {
	Origin          string
	OriginType      PlaceType
	Destination     string
	DestinationType PlaceType
	Weight          int
	Couriers        []string
}
