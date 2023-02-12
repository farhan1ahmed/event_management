package event_tickets

import (
	"event_ticket_service/models"
	"event_ticket_service/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) createEventTicket(ctx *gin.Context){
	logger := utility.GetLogger()
	logger.Info("createEventTicket endpoint called")

	var reqPayload models.CreateEvenTicket
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		logger.Error(err.Error())
		utility.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), true, nil)
		return
	}

	err := server.store.CreateEventTicketInDB(ctx, reqPayload)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := "failed to create event_ticket"
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	msg := "event created successfully"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, nil)
}

func (server *Server) createTicketTypes(ctx *gin.Context){
	logger := utility.GetLogger()
	logger.Info("createTicketTypes endpoint called")

	var reqPayload models.CreateTicketTypes
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		logger.Error(err.Error())
		utility.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), true, nil)
		return
	}

	err := server.store.CreateTicketTypeInDB(ctx, reqPayload)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := fmt.Sprintf("failed to create ticket_type against the event with id %v", reqPayload.EventID)
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	msg := "ticket type created successfully"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, nil)
}

func (server *Server) reserveTicket(ctx *gin.Context){
	logger := utility.GetLogger()
	logger.Info("reserveTicket endpoint called")

	var reqPayload models.ReserveTicket
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		logger.Error(err.Error())
		utility.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), true, nil)
		return
	}

	err := server.store.ReserveTicketInDB(ctx, reqPayload)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := fmt.Sprintf("failed to reserve ticket for the event", reqPayload.EventID)
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	msg := "ticket reserved successfully"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, nil)
}