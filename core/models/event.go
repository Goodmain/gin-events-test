package models

type Event struct {
	BaseModel
	InitialID string  `json:"initial_id"`
	Name      string  `json:"name"`
	URL       string  `json:"url"`
	Users     []*User `json:"users,omitempty" gorm:"many2many:event_users;"`
	City      string  `json:"city"`
}
