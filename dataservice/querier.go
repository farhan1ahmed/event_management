package dataservice

import (
	"context"
	"event_ticket_service/models"
	"github.com/olivere/elastic/v7"
)

type Querier interface{
	InitEventTicketModels() error
	CreateEventTicketInDB(ctx context.Context, reqPayload models.CreateEventTicket) (models.Event, error)
	UpdateEventTicketInDB(ctx context.Context, reqPayload models.UpdateEventTicket) (models.Event, error)
	CreateTicketTypeInDB(ctx context.Context, reqPayload models.CreateTicketTypes) error
	ReserveTicketInDB(ctx context.Context, reqPayload models.ReserveTicket) error
	AddToCartInDB(ctx context.Context, reqPayload models.AddToCart) error
	UpdateCartItemInDB(ctx context.Context, reqPayload models.UpdateCartItem) error
	DeleteCartItemFromDB(ctx context.Context, reqPayload models.DeleteCartItem) error
	GetCartItemsFromDB(ctx context.Context) ([]models.Cart, error)

	//FollowScope
	FollowEventForUser(ctx context.Context, reqPayload models.FollowEvent) error
	FollowEventOrganizerForUser(ctx context.Context, reqPayload models.FollowEventOrganizer) error

	//Notifications
	GetEventUpdateNotificationsFromDB(ctx context.Context, userId int) ([]models.EventChangesNotifications, error)
	GetEventOrganizerNotificationsFromDB(ctx context.Context, userId int) ([]models.EventOrganizerNotifications, error)
	StoreEventDetailsInES(ctx context.Context, ESEvent models.ESEvent) error
	MakeTextBasedSearchESQuery(ctx context.Context, query string, skip int, take int) (*elastic.SearchResult, error)
	MakeLocationBasedSearchESQuery(ctx context.Context, lat float64, long float64, distance string) (*elastic.SearchResult, error)
}

var _ Querier = (*Queries)(nil)