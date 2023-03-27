package elasticsearch

import (
	"context"
	"event_ticket_service/env"
	"event_ticket_service/utility"
	"github.com/olivere/elastic/v7"
)
var elasticClient *elastic.Client

func init(){
	logger := utility.GetLogger()
	logger.Info("Initializing connection with ElasticSearch...")
	eCl, err := elastic.NewClient(elastic.SetURL(env.Env.ElasticURL),
		elastic.SetSniff(false),
	)
	if err != nil{
		logger.Fatal(err.Error())
	}
	logger.Info("Connection successful with ElasticSearch")

	mapping:=
		`{
	  "mappings": {
		"properties": {
		  "name": {
			"type": "text"
		  },
		  "description": {
			"type": "text"
		  },
		  "location": {
			"type": "geo_point"
		  }
		}
	  }
	}`
	ctx := context.Background()
	exists, err := eCl.IndexExists(ElasticIndexName).Do(ctx)
	if err != nil {
		logger.Fatal(err.Error())
	}
	if !exists {
		createIndex, err := eCl.CreateIndex(ElasticIndexName).BodyString(mapping).Do(ctx)
		if err != nil {
			logger.Fatal(err.Error())
		}
		if !createIndex.Acknowledged {
			logger.Fatal("Failed to create ElasticSearch index")
		}
	}

	_, err = eCl.PutMapping().Index(ElasticIndexName).BodyString(mapping).Do(ctx)
	if err != nil {
		// handle error
	}
	elasticClient = eCl
}

func GetElasticClient() *elastic.Client{
	return elasticClient
}

