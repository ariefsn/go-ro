package models

type AwbSummary struct {
	CourierCode  string `json:"courier_code"`
	CourierName  string `json:"courier_name"`
	Awb          string `json:"waybill_number"`
	AwbDate      string `json:"waybill_date"`
	ServiceCode  string `json:"service_code"`
	ShipperName  string `json:"shipper_name"`
	ReceiverName string `json:"receiver_name"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Status       string `json:"status"`
}

type AwbDetails struct {
	Awb              string `json:"waybill_number"`
	AwbDate          string `json:"waybill_date"`
	AwbTime          string `json:"waybill_time"`
	Weight           string `json:"weight"`
	Origin           string `json:"origin"`
	Destination      string `json:"destination"`
	ShipperName      string `json:"shipper_name"`
	ShipperCity      string `json:"shipper_city"`
	ShipperAddress1  string `json:"shipper_address1"`
	ShipperAddress2  string `json:"shipper_address2"`
	ShipperAddress3  string `json:"shipper_address3"`
	ReceiverName     string `json:"receiver_name"`
	ReceiverCity     string `json:"receiver_city"`
	ReceiverAddress1 string `json:"receiver_address1"`
	ReceiverAddress2 string `json:"receiver_address2"`
	ReceiverAddress3 string `json:"receiver_address3"`
}

type DeliveryStatus struct {
	Status      string `json:"status"`
	PodReceiver string `json:"pod_receiver"`
	PodDate     string `json:"pod_date"`
	PodTime     string `json:"pod_time"`
}

type Manifest struct {
	City        string `json:"city_name"`
	Code        int    `json:"manifest_code"`
	Description string `json:"manifest_description"`
	Date        string `json:"manifest_date"`
	Time        string `json:"manifest_time"`
}

type Awb struct {
	Delivered      bool           `json:"delivered"`
	Summary        AwbSummary     `json:"summary"`
	Details        AwbDetails     `json:"details"`
	DeliveryStatus DeliveryStatus `json:"delivery_status"`
	Manifest       []Manifest     `json:"manifest"`
}
