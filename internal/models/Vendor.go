package models

type Vendor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Contact string `json:"contact"`
	Address string `json:"address"`
}
