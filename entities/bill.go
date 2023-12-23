package entities

type Bill struct {
	ID           int64         `json:"id"`
	Code         string        `json:"code"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Host         *User         `json:"host,omitempty"`
	Participants []Participant `json:"participants,omitempty"`
	Items        []Item        `json:"items,omitempty"`
}
