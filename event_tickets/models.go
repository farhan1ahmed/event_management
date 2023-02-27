package event_tickets

import (
	"event_ticket_service/dataservice"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)

type Server struct {
	store       dataservice.Store
	router      *gin.Engine
	routerGroup *gin.RouterGroup
	elastic 	*elastic.Client
}

