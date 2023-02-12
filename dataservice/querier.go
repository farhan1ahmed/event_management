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
}

var _ Querier = (*Queries)(nil)