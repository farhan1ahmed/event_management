package models

import "time"

type Event struct {
	ID                       int       `gorm:"primarykey"`
	EventType                EventType `gorm:"ForeignKey:ID"`
	EventTypeID              int
	EventName                string `gorm:"not null" json:"event_name"`
	EventDescription         string `json:"event_description"`
	EventAddress             string `json:"event_address"`
	EventLocationId          int
	EventCategory            string
	StartTime                *time.Time
	EndTime                  *time.Time
	BookingCloseTime         *time.Time
	EventOrganizer           EventOrganizer `gorm:"ForeignKey:ID"`
	EventOrganizerId         int
	IsDeleted                bool `default:"false"`
	IsCancelled              bool `default:"false"`
	IsSeatManagementRequired bool `default:"false"`
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
	Event             Event `gorm:"ForeignKey:ID"`
	EventID           int
	TicketType        string
	TotalLimit        int
	ReservedQuantity  int `default:"0"`
	RemainingQuantity int
}

type Ticket struct {
	ID            int         `gorm:"primarykey"`
	TicketType    TicketTypes `gorm:"ForeignKey:ID"`
	TicketTypeID  int
	TicketOwner   User `gorm:"ForeignKey:ID"`
	TicketOwnerID int
	IsPaid        bool `default:"false"`
	IsActive      bool
}
