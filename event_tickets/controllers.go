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
	rg.Handle(http.MethodPost, "event-ticket-types", server.createTicketTypes)
	rg.Handle(http.MethodPost, "ticket", server.reserveTicket)
}

