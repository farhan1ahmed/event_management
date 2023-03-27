package main

import (
	"crypto/tls"
	"event_ticket_service/dataservice"
	"event_ticket_service/elasticsearch"
	"event_ticket_service/env"
	"event_ticket_service/event_tickets"
	"event_ticket_service/utility"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	logger := utility.GetLogger()
	db := dataservice.GetDBConnection()
	elasticClient := elasticsearch.GetElasticClient()

	r := gin.Default()
	rg := r.Group("api/v1/")

	dbStore := dataservice.NewStore(db, elasticClient)

	event_tickets.CreateNewServer(dbStore, r, rg)

	err := r.Run(env.Env.GetServerAddress())
	if err != nil {
		logger.Error(err.Error())
		_ = logger.Sync()
		return
	}
}
