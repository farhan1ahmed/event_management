package models

import (
	"time"
)

type Cart struct {
	ID           int `gorm:"primarykey"`
	TicketTypeID int
	TicketType TicketTypes `gorm:"ForeignKey:ID"`
	UserID       int
	User         User `gorm:"ForeignKey:ID"`
	Quantity     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
