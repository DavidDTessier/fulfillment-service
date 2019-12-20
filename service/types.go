package service

type fulfillmentStatus struct {
	SKU             string `json:"sku"`
	ShipsWithin     int    `json:"shipsWithin"`
	QuantityInStock int    `json:"quantityInStock"`
}
