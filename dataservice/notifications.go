package dataservice

import (
	"context"
	"event_ticket_service/models"
	"log"
)

func (q *Queries)GetEventUpdateNotificationsFromDB(ctx context.Context, userId int) ([]models.EventChangesNotifications, error){
	var notification []models.EventChangesNotifications
	result := q.db.Find(&notification,"user_id = ?", userId)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}
	for i := range notification{
		var noti models.Notification
		result = q.db.First(&noti,"id = ?", notification[i].NotificationID)
		if result.Error != nil{
			log.Println(result.Error)
			return nil, result.Error
		}
		notification[i].Notification = noti
	}
	return notification, nil
}

func (q *Queries)GetEventOrganizerNotificationsFromDB(ctx context.Context, userId int) ([]models.EventOrganizerNotifications, error){
	var notification []models.EventOrganizerNotifications
	result := q.db.Find(&notification,"user_id = ?", userId)
	if result.Error != nil{
		log.Println(result.Error)
		return nil, result.Error
	}
	for i := range notification{
		var noti models.Notification
		result = q.db.First(&noti,"id = ?", notification[i].NotificationID)
		if result.Error != nil{
			log.Println(result.Error)
			return nil, result.Error
		}
		notification[i].Notification = noti
	}
	return notification, nil
}
