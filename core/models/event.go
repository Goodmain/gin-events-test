package models

import (
	"time"
)

type Event struct {
	BaseModel
	InitialID   string    `json:"initial_id"`
	Name        string    `json:"name"`
	Description string    `json:"text"`
	URL         string    `json:"url"`
	Users       []*User   `json:"users,omitempty" gorm:"many2many:event_users;"`
	City        string    `json:"city"`
	Address     string    `json:"address"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Zip         string    `json:"zip"`
	Place       string    `json:"place"`
	Image       string    `json:"image"`
	Thumbnail   string    `json:"thumbnail"`
	Date        time.Time `json:"date"`
	Lng         string    `json:"lng"`
	Lat         string    `json:"lat"`
}
