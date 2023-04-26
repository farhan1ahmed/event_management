package dataservice

import (
	"context"
	"event_ticket_service/models"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"time"
)

func (q *Queries)CreateEventTicketInDB(ctx context.Context, reqPayload models.CreateEventTicket) (models.Event, error){
	var event models.Event
	err := q.db.Transaction(func(tx *gorm.DB) error {
		// Create event organizer
		eventOrganizer := models.EventOrganizer{
			OrganizerName: reqPayload.OrganizerName,
			OrganizerAddress: reqPayload.OrganizerAddress,
			OrganizerContact: reqPayload.OrganizerContact,
		}
		result := tx.Model(&eventOrganizer).First(&eventOrganizer, "organizer_name = ?", eventOrganizer.OrganizerName)
		if eventOrganizer.ID == 0 {
			result = tx.Model(&eventOrganizer).Create(&eventOrganizer)
			if result.Error != nil {
				log.Println(result.Error)
				return result.Error
			}
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
		event = models.Event{
			EventTypeID: reqPayload.EventTypeID,
			EventName: reqPayload.EventName,
			EventAddress: reqPayload.EventAddress,
			EventDescription: reqPayload.EventDescription,
			EventCategory: reqPayload.EventCategory,
			StartTime: reqPayload.StartTime,
			EndTime: reqPayload.EndTime,
			BookingCloseTime: &bookingCloseTime,
			EventOrganizerId: eventOrganizer.ID,
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
			EventID: event.ID,
		}
		result = tx.Model(&ticketTypes).Create(&ticketTypes)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}

		// Get EventOrganizer Followers
		var eventOrgFollowers []models.EventOrganizerFollowers
		result = q.db.Find(&eventOrgFollowers, "event_organizer_id = ?", eventOrganizer.ID)
		if result.Error != nil {
			log.Println(result.Error)
			return result.Error
		}

		if len(eventOrgFollowers) != 0 {
			//capture events for event_organizer_notification table
			var notificationMessage string
			notificationMessage = fmt.Sprintf("Event organizer %s published a new event %s", eventOrganizer.OrganizerName, event.EventName)

			var ntfn models.Notification
			ntfn.Updates = notificationMessage
			result = q.db.Model(&ntfn).Create(&ntfn)
			if result.Error != nil {
				log.Println(result.Error)
				return result.Error
			}

			for _, follower := range eventOrgFollowers {
				var eventNtfn models.EventOrganizerNotifications
				eventNtfn.EventOrganizerID = eventOrganizer.ID
				eventNtfn.UserID = follower.UserID
				eventNtfn.NotificationID = ntfn.ID
				result = q.db.Model(&eventNtfn).Create(&eventNtfn)
				if result.Error != nil {
					log.Println(result.Error)
					return result.Error
				}
				}
		}
		return nil
	})
	if err != nil{
		log.Println(err.Error())
		return event, err
	}
	return event, nil
}

func (q *Queries)UpdateEventTicketInDB(ctx context.Context, reqPayload models.UpdateEventTicket) (models.Event, error){
	var updatedEventDetails models.Event
	err := q.db.Transaction(func(tx *gorm.DB) error {
		var oldEventDetails models.Event
		updatedEventDetails = models.Event{
			ID:               reqPayload.EventID,
			EventName:        reqPayload.EventName,
			EventAddress:     reqPayload.EventAddress,
			EventDescription: reqPayload.EventDescription,
			StartTime:        reqPayload.StartTime,
			EndTime:          reqPayload.EndTime,
		}

		result := q.db.First(&oldEventDetails, reqPayload.EventID)
		if result.Error != nil {
			log.Println(result.Error)
			return result.Error
		}
		result = q.db.Model(&updatedEventDetails).Updates(&updatedEventDetails)
		if result.Error != nil {
			log.Println(result.Error)
			return result.Error
		}
		// Get Event Followers
		var eventFollowers []models.EventFollowers
		result = q.db.Find(&eventFollowers, "event_id = ?", reqPayload.EventID)
		if result.Error != nil {
			log.Println(result.Error)
			return result.Error
		}

		if len(eventFollowers) != 0 {
			//capture change in event_changes_notification table
			var notifications []string
			var notificationMessage string
			if oldEventDetails.EventName != updatedEventDetails.EventName {
				notificationMessage = fmt.Sprintf("Name for %s event has been updated to %s. ", oldEventDetails.EventName, updatedEventDetails.EventName)
				notifications = append(notifications, notificationMessage)
			}
			if oldEventDetails.EventAddress != updatedEventDetails.EventAddress {
				notificationMessage = fmt.Sprintf("Address for %s event has been updated to %s. ", oldEventDetails.EventName, updatedEventDetails.EventAddress)
				notifications = append(notifications, notificationMessage)
			}

			if oldEventDetails.EventDescription != updatedEventDetails.EventDescription {
				notificationMessage = fmt.Sprintf("Description for %s event has been updated to %s. ", oldEventDetails.EventName, updatedEventDetails.EventDescription)
				notifications = append(notifications, notificationMessage)
			}

			for _, message := range notifications {
				var ntfn models.Notification
				ntfn.Updates = message
				result = q.db.Model(&ntfn).Create(&ntfn)
				if result.Error != nil {
					log.Println(result.Error)
					return result.Error
				}

				for _, follower := range eventFollowers {
					var eventNtfn models.EventChangesNotifications
					eventNtfn.EventID = reqPayload.EventID
					eventNtfn.UserID = follower.UserID
					eventNtfn.NotificationID = ntfn.ID
					result = q.db.Model(&eventNtfn).Create(&eventNtfn)
					if result.Error != nil {
						log.Println(result.Error)
						return result.Error
					}
				}
			}
		}
		return nil
	})
	if err != nil{
		log.Println(err.Error())
		return updatedEventDetails, err
	}
	return updatedEventDetails, nil
}

func (q *Queries)CreateTicketTypeInDB(ctx context.Context, reqPayload models.CreateTicketTypes) error{
	ticketType := models.TicketTypes{
		EventID: reqPayload.EventID,
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
		if ticketType.RemainingQuantity >0 {
			ticketType.ReservedQuantity  = ticketType.ReservedQuantity + 1
			ticketType.RemainingQuantity  = ticketType.RemainingQuantity - 1
		}
		result = tx.Model(&ticketType).Updates(ticketType)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}

		// Get event details
		var event models.Event
		event.ID = reqPayload.EventID
		result = tx.First(&event)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}

		// Create Ticket
		ticketType.EventID = event.ID
		ticket := models.Ticket{
			TicketTypeID: ticketType.ID,
			TicketOwnerID: reqPayload.TicketOwnerID,
			IsPaid: false,
			IsActive: false,
		}

		result = tx.Create(&ticket)
		if result.Error != nil{
			log.Println(result.Error)
			return result.Error
		}

		// Create seat reservation if seat management enabled
		if event.IsSeatManagementRequired{
			// Get seat
			var seat models.Seat
			var reservedSeat models.ReservedSeat
			var exists bool
			seat.ID = reqPayload.SeatID
			reservedSeat.SeatID = reqPayload.SeatID
			// Check if seat not reserved already
			result = tx.Model(reservedSeat).
				Select("id").
				Where("seat_id = ?", reqPayload.SeatID).
				Find(&exists)
			if result.Error != nil{
				return result.Error
			}
			if exists{
				return errors.New("seat reserved already")
			}


			result = tx.First(&seat)
			if result.Error != nil{
				log.Println(result.Error)
				return result.Error
			}
			reservedSeat = models.ReservedSeat{
				TicketID: ticket.ID,
				EventID: event.ID,
				SeatID: seat.ID,
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