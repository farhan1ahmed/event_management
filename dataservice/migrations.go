package dataservice

import (
	"event_ticket_service/models"
	"gorm.io/gorm"
)

func (q *Queries)InitEventTicketModels() error{
	err := q.db.AutoMigrate(&models.EventType{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.EventOrganizer{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.Event{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.TicketTypes{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.Cart{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.User{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.Ticket{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.Seat{})
	if err!= nil{
		return err
	}
	err = q.db.AutoMigrate(&models.ReservedSeat{})
	if err!= nil{
		return err
	}

	var eventTypes []models.EventType
	eventTypes = []models.EventType{
		{1, "physical"},{2, "online"},{3, "hybrid"},
	}

	var seats []models.Seat
	var seat models.Seat
	id := 0
	for i:=0;i<10;i++ {
		for j:=0;j<10;j++{
			id = id+1;
			seat = models.Seat{ID:id, Row: i, Number: j}
			seats = append(seats, seat)
		}
	}
	eventTypes = []models.EventType{
		{1, "physical"},{2, "online"},{3, "hybrid"},
	}
	err = q.db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&eventTypes).Create(eventTypes)
	return nil
	})
	err = q.db.Transaction(func(tx *gorm.DB) error {
		tx.Model(&seats).Create(seats)
		return nil
	})
	return nil
}
