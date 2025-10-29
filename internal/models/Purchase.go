package models

type Purchase struct {
	ID          int     `json:"id"`
	VendorID    int     `json:"vendor_id"`
	Status      string  `json:"status"`
	TotalAmount float64 `json:"total_amount"`
}

type PurchaseDetail struct {
	ID            int     `json:"id"`
	PurchaseID    int     `json:"order_id"`
	MakhanaTypeID int     `json:"makhana_type_id"`
	Quantity      float64 `json:"quantity"`
	Price         float64 `json:"price"`
	LineItemTotal float64 `json:"line_item_total"`
}
