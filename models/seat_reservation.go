package models

type Seat struct {
	ID     int `gorm:"primarykey"`
	Row    int
	Number int
}

type ReservedSeat struct {
	ID       int    `gorm:"primarykey"`
	Ticket   Ticket `gorm:"ForeignKey:ID"`
	Seat     Seat   `gorm:"ForeignKey:ID"`
	Event    Event  `gorm:"ForeignKey:ID"`
	TicketID int
	SeatID   int
	EventID  int
}
