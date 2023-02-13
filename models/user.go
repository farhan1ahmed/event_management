package models

type User struct {
	ID       int    `gorm:"primarykey"`
	Username string
	FirstName string
	LastName string
	Address string
}
