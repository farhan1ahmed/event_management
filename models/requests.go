package models

import "time"

type CreateEventTicket struct {
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

type UpdateEventTicket struct {
	EventID          int        `json:"event_id"`
	EventName        string     `json:"event_name"`
	EventDescription string     `json:"event_description"`
	EventAddress     string     `json:"event_address"`
	EventLocationId  int        `json:"event_location_id"`
	StartTime        *time.Time `json:"start_time"`
	EndTime          *time.Time `json:"end_time"`
	BookingCloseTime *time.Time `json:"booking_close_time"`
}

type CreateTicketTypes struct {
	EventID    int    `json:"event_id"`
	TicketType string `json:"ticket_type"`
	TotalLimit int    `json:"total_limit"`
}

type ReserveTicket struct {
	EventID       int `json:"event_id"`
	TicketTypeID  int `json:"ticket_type_id"`
	TicketOwnerID int `json:"ticket_owner_id"`
	SeatID        int `json:"seat_id"`
}

type AddToCart struct {
	UserID       int `json:"user_id"`
	TicketTypeID int `json:"ticket_type_id"`
	Quantity     int `json:"quantity"`
}

type UpdateCartItem struct {
	CartItemID int `json:"cart_item_id"`
	Quantity   int `json:"quantity"`
}

type DeleteCartItem struct {
	CartItemID int `json:"cart_item_id"`
}

type FollowEvent struct {
	EventID int `json:"event_id"`
	UserID  int `json:"user_id"`
}

type FollowEventOrganizer struct {
	EventOrganizerID int `json:"event_organizer_id"`
	UserID           int `json:"user_id"`
}
