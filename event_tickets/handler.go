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
		msg := fmt.Sprintf("failed to reserve ticket for the event %v", reqPayload.EventID)
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	msg := "ticket reserved successfully"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, nil)
}

func (server *Server) addToCart(ctx *gin.Context){
	logger := utility.GetLogger()
	logger.Info("addToCart endpoint called")

	var reqPayload models.AddToCart
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		logger.Error(err.Error())
		utility.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), true, nil)
		return
	}

	err := server.store.AddToCartInDB(ctx, reqPayload)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := "failed to add item to cart"
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	msg := "item added to cart successfully"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, nil)
}

func (server *Server) updateCartItem(ctx *gin.Context){
	logger := utility.GetLogger()
	logger.Info("updateCartItem endpoint called")

	var reqPayload models.UpdateCartItem
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		logger.Error(err.Error())
		utility.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), true, nil)
		return
	}

	err := server.store.UpdateCartItemInDB(ctx, reqPayload)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := "failed to update item in cart"
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	msg := "item updated in cart successfully"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, nil)
}

func (server *Server) deleteCartItem(ctx *gin.Context){
	logger := utility.GetLogger()
	logger.Info("updateCartItem endpoint called")

	var reqPayload models.DeleteCartItem
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		logger.Error(err.Error())
		utility.GenerateResponse(ctx, http.StatusBadRequest, err.Error(), true, nil)
		return
	}

	err := server.store.DeleteCartItemFromDB(ctx, reqPayload)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := "failed to delete item from cart"
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	msg := "cart item deleted successfully"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, nil)
}


func (server *Server) showCartItems(ctx *gin.Context){
	logger := utility.GetLogger()
	logger.Info("showCartItems endpoint called")

	cartItems, err := server.store.GetCartItemsFromDB(ctx)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := "failed to gets items from cart"
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	var cart models.CartItems
	var cartItemsResponse []models.CartItems
	for _, item := range cartItems{
		cart.ID = item.ID
		cart.TicketTypeID = item.TicketTypeID
		cart.UserID = item.UserID
		cart.Quantity = item.Quantity
		cart.CreatedAt = item.CreatedAt
		cart.UpdatedAt = item.UpdatedAt
		cartItemsResponse = append(cartItemsResponse, cart)
	}
	msg := "successfully retreived cart items"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, cartItemsResponse)
}