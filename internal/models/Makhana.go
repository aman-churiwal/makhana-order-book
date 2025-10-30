package models

type Makhana struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Rate        float64 `json:"rate"`
	VendorID    int     `json:"vendor_id"`
}
