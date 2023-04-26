package event_tickets

import (
	"event_ticket_service/dataservice"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateNewServer(dbStore dataservice.Store, router *gin.Engine, rg *gin.RouterGroup){
	server := &Server{
		store: dbStore,
		router: router,
		routerGroup: rg,
	}
	err := server.store.InitEventTicketModels()
	if err != nil{
		log.Println(err.Error())
	}
	rg.Handle(http.MethodPost, "event-ticket", server.createEventTicket)
	rg.Handle(http.MethodPatch, "event-ticket", server.updateEventTicket)
	rg.Handle(http.MethodPost, "event-ticket-types", server.createTicketTypes)
	rg.Handle(http.MethodPost, "ticket", server.reserveTicket)
	rg.Handle(http.MethodPost, "cart", server.addToCart)
	rg.Handle(http.MethodPatch, "cart", server.updateCartItem)
	rg.Handle(http.MethodDelete, "cart", server.deleteCartItem)
	rg.Handle(http.MethodGet, "admin/cart", server.showCartItems)
	rg.Handle(http.MethodGet, "search", server.searchEngine)

	//FollowScope
	rg.Handle(http.MethodPost, "follow/event", server.followEvent)
	rg.Handle(http.MethodPost, "follow/event-organizer", server.followEventOrganizer)

	// Notifications
	rg.Handle(http.MethodGet, "event-notifications/:user_id", server.eventUpdateNotifications)
	rg.Handle(http.MethodGet, "event-organizer-notifications/:user_id", server.eventOrganizerNotifications)
}

