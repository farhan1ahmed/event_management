package elasticsearch

import (
	"event_ticket_service/env"
	"event_ticket_service/utility"
	"github.com/olivere/elastic/v7"
)
var elasticClient *elastic.Client

func init(){
	logger := utility.GetLogger()
	eCl, err := elastic.NewClient(elastic.SetURL(env.Env.ElasticURL),
		elastic.SetSniff(false),
	)
	if err != nil{
		logger.Error(err.Error())
	}
	elasticClient = eCl
}

func GetElasticClient() *elastic.Client{
	return elasticClient
}

