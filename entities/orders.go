package entities

type Order struct {
	ID          int64        `json:"id"`
	Qty         int64        `json:"qty"`
	Note        string       `json:"note"`
	Item        *Item        `json:"item,omitempty"`
	Participant *Participant `json:"participant,omitempty"`
	Subtotal    float64      `json:"subtotal"`
}
