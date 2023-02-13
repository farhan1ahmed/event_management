package dataservice

import (
	"context"
	"event_ticket_service/models"
)

type Querier interface{
	InitEventTicketModels() error
	CreateEventTicketInDB(ctx context.Context, reqPayload models.CreateEvenTicket) error
	CreateTicketTypeInDB(ctx context.Context, reqPayload models.CreateTicketTypes) error
	ReserveTicketInDB(ctx context.Context, reqPayload models.ReserveTicket) error
	AddToCartInDB(ctx context.Context, reqPayload models.AddToCart) error
	UpdateCartItemInDB(ctx context.Context, reqPayload models.UpdateCartItem) error
	DeleteCartItemFromDB(ctx context.Context, reqPayload models.DeleteCartItem) error
	GetCartItemsFromDB(ctx context.Context) ([]models.Cart, error)
}

var _ Querier = (*Queries)(nil)