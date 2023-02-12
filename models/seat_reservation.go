package models

type Seat struct {
	ID     int `gorm:"primarykey"`
	Row    int
	Number int
}

type ReservedSeat struct {
	ID       int  `gorm:"primarykey"`
	TicketID Ticket  `gorm:"ForeignKey:ID"`
	SeatID   Seat `gorm:"ForeignKey:ID"`
	EventID  Event  `gorm:"ForeignKey:ID"`
}
