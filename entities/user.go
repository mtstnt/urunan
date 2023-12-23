package entities

type User struct {
	ID             int64         `json:"id" db:"id"`
	Email          string        `json:"email" db:"email"`
	FullName       string        `json:"full_name" db:"full_name"`
	Password       string        `json:"-" db:"password"`
	HostedBills    []Bill        `json:"hosted_bills,omitempty"`
	Participations []Participant `json:"participations,omitempty"`
}
