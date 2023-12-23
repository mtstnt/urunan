package entities

type Item struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	InitialQty int64   `json:"initial_qty"`
	Bill       *Bill   `json:"bill,omitempty"`
}
