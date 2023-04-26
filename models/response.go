package models

import "time"

type CartItems struct {
	ID           int `gorm:"primarykey"`
	TicketTypeID int
	UserID       int
	Quantity     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type NotificationResponse struct {
	Status string `gorm:"status" json:"status"`
	Notification
}