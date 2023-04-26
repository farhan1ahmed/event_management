package models

type Notification struct {
	ID      int `gorm:"primarykey" json:"id"`
	Updates string `gorm:"updates" json:"update"`
}
type EventChangesNotifications struct {
	ID             int `gorm:"primarykey"`
	Event          Event `gorm:"ForeignKey:ID"`
	EventID        int
	User           User `gorm:"ForeignKey:ID"`
	UserID         int
	Notification   Notification `gorm:"ForeignKey:ID"`
	NotificationID int
	Status         string       `gorm:"default:'unread'"`
}

type EventOrganizerNotifications struct {
	ID               int `gorm:"primarykey"`
	EventOrganizerID int
	EventOrganizer   EventOrganizer `gorm:"ForeignKey:ID"`
	UserID           int
	NotificationID   int
	Notification     Notification `gorm:"ForeignKey:ID"`
	User             User         `gorm:"ForeignKey:ID"`
	Status           string       `gorm:"default:'unread'"`
}

type EventFollowers struct {
	ID      int `gorm:"primarykey"`
	EventID int
	Event   Event `gorm:"ForeignKey:ID"`
	UserID  int
	User    User `gorm:"ForeignKey:ID"`
}

type EventOrganizerFollowers struct {
	ID               int `gorm:"primarykey"`
	EventOrganizerID int
	EventOrganizer   EventOrganizer `gorm:"ForeignKey:ID"`
	UserID           int
	User             User `gorm:"ForeignKey:ID"`
}
