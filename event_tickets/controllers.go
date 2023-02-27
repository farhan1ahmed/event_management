package event_tickets

import (
	"event_ticket_service/dataservice"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
)

func CreateNewServer(dbStore dataservice.Store, router *gin.Engine, rg *gin.RouterGroup, eCl *elastic.Client){
	server := &Server{
		store: dbStore,
		router: router,
		routerGroup: rg,
		elastic: eCl,
	}
	err := server.store.InitEventTicketModels()
	if err != nil{
		log.Println(err.Error())
	}
	rg.Handle(http.MethodPost, "event-ticket", server.createEventTicket)
	rg.Handle(http.MethodPost, "event-ticket-types", server.createTicketTypes)
	rg.Handle(http.MethodPost, "ticket", server.reserveTicket)
	rg.Handle(http.MethodPost, "cart", server.addToCart)
	rg.Handle(http.MethodPatch, "cart", server.updateCartItem)
	rg.Handle(http.MethodDelete, "cart", server.deleteCartItem)
	rg.Handle(http.MethodGet, "admin/cart", server.showCartItems)
	rg.Handle(http.MethodGet, "search", server.searchEngine)
}

