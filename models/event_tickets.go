package models

import "time"

type Event struct {
	ID                       int       `gorm:"primarykey"`
	EventTypeID              EventType `gorm:"ForeignKey:ID"`
	EventName                string    `gorm:"not null"`
	EventDescription         string
	EventAddress             string
	EventLocationId          int
	EventCategory            string
	StartTime                *time.Time
	EndTime                  *time.Time
	BookingCloseTime         *time.Time
	EventOrganizerId         EventOrganizer `gorm:"ForeignKey:ID"`
	IsDeleted                bool           `default:"false"`
	IsCancelled              bool           `default:"false"`
	IsSeatManagementRequired bool           `default:"false"`
}

type EventType struct {
	ID            int    `gorm:"primarykey"`
	EventTypeName string `gorm:"not null"`
}

type EventOrganizer struct {
	ID               int    `gorm:"primarykey"`
	OrganizerName    string `gorm:"not null"`
	OrganizerContact string
	OrganizerAddress string
}

type TicketTypes struct {
	ID                int   `gorm:"primarykey"`
	EventID           Event `gorm:"ForeignKey:ID"`
	TicketType        string
	TotalLimit        int
	ReservedQuantity  int `default:"0"`
	RemainingQuantity int
}

type Ticket struct {
	ID                 int         `gorm:"primarykey"`
	TicketTypeID       TicketTypes `gorm:"ForeignKey:ID"`
	TicketOwnerContact string
	IsPaid             bool `default:"false"`
	IsActive           bool
}
