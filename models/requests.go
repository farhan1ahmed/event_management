package models

import "time"

type CreateEvenTicket struct {
	EventTypeID              int        `json:"event_type_id"`
	EventName                string     `json:"event_name"`
	EventDescription         string     `json:"event_description"`
	EventAddress             string     `json:"event_address"`
	EventLocationId          int        `json:"event_location_id"`
	EventCategory            string     `json:"event_category"`
	StartTime                *time.Time `json:"start_time"`
	EndTime                  *time.Time `json:"end_time"`
	BookingCloseTime         *time.Time `json:"booking_close_time"`
	IsSeatManagementRequired bool       `json:"is_seat_management_required"`
	OrganizerName            string     `json:"organizer_name"`
	OrganizerContact         string     `json:"organizer_contact"`
	OrganizerAddress         string     `json:"organizer_address"`
	TicketType               string     `json:"ticket_type"`
	TotalLimit               int        `json:"total_limit"`
}

type CreateTicketTypes struct {
	EventID             int    `json:"event_id"`
	TicketType        string `json:"ticket_type"`
	TotalLimit        int    `json:"total_limit"`
}

type ReserveTicket struct {
	EventID             int    `json:"event_id"`
	TicketTypeID        int `json:"ticket_type_id"`
	TicketOwnerContact  string `json:"ticket_owner_contact"`
	SeatID				int 	`json:"seat_id"`
}