package entities

type Participant struct {
	ID       int64  `json:"id"`
	Bill     *Bill  `json:"bill,omitempty"`
	User     *User  `json:"user,omitempty"`
	Nickname string `json:"nickname"`

	Orders []Order `json:"orders"`
}
