package dataservice

import (
	"context"
	"event_ticket_service/models"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

func (q *Queries)CreateEventTicketInDB(ctx context.Context, reqPayload models.CreateEvenTicket) error{
	err := q.db.Transaction(func(tx *gorm.DB) error {
		// Create event organizer
		eventOrganizer := models.EventOrganizer{
			OrganizerName: reqPayload.OrganizerName,
			OrganizerAddress: reqPayload.OrganizerAddress,
			OrganizerContact: reqPayload.OrganizerContact,
		}
		result := tx.Model(&eventOrganizer).Create(&eventOrganizer)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}

		// set bookingClose time to time specified in request body
		// otherwise set it to event start time
		var bookingCloseTime time.Time
		if reqPayload.BookingCloseTime != nil{
			bookingCloseTime = *reqPayload.BookingCloseTime
		}else{
			bookingCloseTime = *reqPayload.StartTime
		}
		// Create event
		event := models.Event{
			EventTypeID: models.EventType{ID:reqPayload.EventTypeID},
			EventName: reqPayload.EventName,
			EventAddress: reqPayload.EventAddress,
			EventDescription: reqPayload.EventDescription,
			EventCategory: reqPayload.EventCategory,
			StartTime: reqPayload.StartTime,
			EndTime: reqPayload.EndTime,
			BookingCloseTime: &bookingCloseTime,
			EventOrganizerId: eventOrganizer,
			IsDeleted: false,
			IsCancelled: false,
			IsSeatManagementRequired: reqPayload.IsSeatManagementRequired,
		}
		result = tx.Model(&event).Create(&event)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}

		// Create ticket types for the event
		ticketTypes := models.TicketTypes{
			TicketType: reqPayload.TicketType,
			TotalLimit: reqPayload.TotalLimit,
			RemainingQuantity: reqPayload.TotalLimit,
			EventID: event,
		}
		result = tx.Model(&ticketTypes).Create(&ticketTypes)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}

		return nil
	})
	if err != nil{
		log.Println(err.Error())
		return err
	}
	return nil
}

func (q *Queries)CreateTicketTypeInDB(ctx context.Context, reqPayload models.CreateTicketTypes) error{
	ticketType := models.TicketTypes{
		EventID: models.Event{ID: reqPayload.EventID},
		TicketType: reqPayload.TicketType,
		TotalLimit: reqPayload.TotalLimit,
		ReservedQuantity: 0,
		RemainingQuantity: reqPayload.TotalLimit,
	}
	result := q.db.Model(&ticketType).Create(&ticketType)
	if result.Error != nil{
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (q *Queries)ReserveTicketInDB(ctx context.Context, reqPayload models.ReserveTicket) error{

	err := q.db.Transaction(func(tx *gorm.DB) error {
		// Check if there's limit remaining for the ticket type
		ticketTypeID := reqPayload.TicketTypeID
		var ticketType models.TicketTypes
		ticketType.ID = ticketTypeID
		result := tx.First(&ticketType)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}
		fmt.Println("1")
		if ticketType.RemainingQuantity >0 {
			ticketType.ReservedQuantity  = ticketType.ReservedQuantity + 1
			ticketType.RemainingQuantity  = ticketType.RemainingQuantity - 1
		}
		result = tx.Model(&ticketType).Updates(ticketType)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}
		fmt.Println("2")

		// Get event details
		var event models.Event
		event.ID = reqPayload.EventID
		result = tx.First(&event)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}
		fmt.Println(event)

		// Create Ticket
		ticketType.EventID = event
		ticket := models.Ticket{
			TicketTypeID: ticketType,
			TicketOwnerContact: reqPayload.TicketOwnerContact,
			IsPaid: false,
			IsActive: false,
		}
		fmt.Println("3")
		fmt.Println(ticket)
		result = tx.Create(&ticket)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}
		fmt.Println("4")
		fmt.Println("5")
		// Create seat reservation if seat management enabled
		if event.IsSeatManagementRequired{
			// Get seat
			var seat models.Seat
			seat.ID = reqPayload.SeatID
			result = tx.First(&seat)
			if result.Error != nil{
				log.Println(result.Error)
				return result.Error
			}
			reservedSeat := models.ReservedSeat{
				TicketID: ticket,
				EventID: event,
				SeatID: seat,
			}
			result = tx.Create(&reservedSeat)
			if result.Error != nil{
				log.Println(result.Error)
				return result.Error
			}
		}

		return nil
	})
	if err != nil{
		log.Println(err.Error())
		return err
	}
	return nil
}