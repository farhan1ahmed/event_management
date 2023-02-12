package event_tickets

import (
	"event_ticket_service/dataservice"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store       dataservice.Store
	router      *gin.Engine
	routerGroup *gin.RouterGroup
}

