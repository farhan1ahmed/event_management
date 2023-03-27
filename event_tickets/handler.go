package event_tickets

import (
	"encoding/json"
	"event_ticket_service/models"
	"event_ticket_service/utility"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"net/http"
	"net/url"
	"strconv"
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

	event, err := server.store.CreateEventTicketInDB(ctx, reqPayload)
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := "failed to create event_ticket"
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	var ESEvent models.ESEvent
	ESEvent.Event = event
	lat, long, err := utility.GetLatLon(url.QueryEscape(event.EventAddress))
	if err != nil{
		logger.Error(err.Error())
		statusCode := http.StatusInternalServerError
		msg := "failed to create event_ticket"
		utility.GenerateResponse(ctx, statusCode, msg, true, nil)
		return
	}

	ESEvent.Location.Latitude = lat
	ESEvent.Location.Longitude = long

	err = server.store.StoreEventDetailsInES(ctx, ESEvent)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to insert event to ES: %v", err))
		msg := "failed to create event_ticket"
		utility.GenerateResponse(ctx, http.StatusInternalServerError, msg, true, nil)
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

func (server *Server) searchEngine(ctx *gin.Context) {
	logger := utility.GetLogger()
	logger.Info("searchEngine endpoint called")

	query := ctx.Query("query")
	if query == "" {
		msg := "Query not specified"
		utility.GenerateResponse(ctx, http.StatusBadRequest, msg, true, nil)
		return
	}
	skip := 0
	take := 10
	if i, err := strconv.Atoi(ctx.Query("skip")); err == nil {
		skip = i
	}
	if i, err := strconv.Atoi(ctx.Query("take")); err == nil {
		take = i
	}

	searchType := ctx.Query("type")
	var result *elastic.SearchResult

	switch searchType {
	case "text":
		_result, err := server.store.MakeTextBasedSearchESQuery(ctx, query, skip, take)
		if err != nil {
			logger.Error(err.Error())
			msg := "Something went wrong"
			utility.GenerateResponse(ctx, http.StatusInternalServerError, msg, true, nil)
			return
		}
		result = _result
	case "location":
		lat, long, err := utility.GetLatLon(url.QueryEscape(query))
		if err != nil {
			logger.Error(err.Error())
			statusCode := http.StatusInternalServerError
			msg := "failed to retrieve results for the search query"
			utility.GenerateResponse(ctx, statusCode, msg, true, nil)
			return
		}

		distance := ctx.Query("distance")
		if distance == ""{
			distance = "10km"
		}
		_result, err := server.store.MakeLocationBasedSearchESQuery(ctx, lat, long, distance)

		if err != nil {
			logger.Error(err.Error())
			msg := "Something went wrong"
			utility.GenerateResponse(ctx, http.StatusInternalServerError, msg, true, nil)
			return
		}
		result = _result
	default:
		msg := "Search type missing"
		utility.GenerateResponse(ctx, http.StatusBadRequest, msg, true, nil)
		return
	}

	res := models.SearchResponse{
		Time: fmt.Sprintf("%d", result.TookInMillis),
		Hits: fmt.Sprintf("%d", result.Hits.TotalHits.Value),
	}
	docs := make([]models.DocumentResponse, 0)
	for _, hit := range result.Hits.Hits {
		var doc models.DocumentResponse
		json.Unmarshal(hit.Source, &doc)
		docs = append(docs, doc)
	}
	res.ResultDocuments = docs
	msg := "successfully retrieved search items"
	utility.GenerateResponse(ctx, http.StatusOK, msg, false, res)
}