package dataservice

import (
	"context"
	"event_ticket_service/models"
	"log"
)

func (q *Queries)FollowEventForUser(ctx context.Context, reqPayload models.FollowEvent) error {
	eventFollower := models.EventFollowers{
		EventID: reqPayload.EventID,
		UserID: reqPayload.UserID,
	}
	result := q.db.Model(&eventFollower).Create(&eventFollower)
	if result.Error != nil{
		log.Println(result.Error)
		return result.Error
	}
	return nil
}

func (q *Queries)FollowEventOrganizerForUser(ctx context.Context, reqPayload models.FollowEventOrganizer) error {
	eventOrganizerFollower := models.EventOrganizerFollowers{
		EventOrganizerID: reqPayload.EventOrganizerID,
		UserID: reqPayload.UserID,
	}
	result := q.db.Model(&eventOrganizerFollower).Create(&eventOrganizerFollower)
	if result.Error != nil{
		log.Println(result.Error)
		return result.Error
	}
	return nil
	return nil
}
