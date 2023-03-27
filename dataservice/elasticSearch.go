package dataservice

import (
	"context"
	"event_ticket_service/elasticsearch"
	"event_ticket_service/models"
	"github.com/olivere/elastic/v7"
)

func (q *Queries) StoreEventDetailsInES(ctx context.Context, ESEvent models.ESEvent) error{
	_, err := q.es.Index().
		Index(elasticsearch.ElasticIndexName).
		BodyJson(ESEvent).
		Do(ctx)
	if err != nil{
		return err
	}
	return nil
}

func (q *Queries) MakeTextBasedSearchESQuery(ctx context.Context, query string, skip int, take int) (*elastic.SearchResult, error){
	esQuery := elastic.NewMultiMatchQuery(query, "event_name", "event_description")
	result, err := q.es.Search().
		Index(elasticsearch.ElasticIndexName).
		Query(esQuery).From(skip).Size(take).Do(ctx)
	if err != nil{
		return nil, err
	}

	return result, nil
}

func (q *Queries) MakeLocationBasedSearchESQuery(ctx context.Context, lat float64, long float64, distance string) (*elastic.SearchResult, error){
	esQuery := elastic.NewGeoDistanceQuery(elasticsearch.ElasticLocationName).Lat(lat).Lon(long).Distance(distance)
	result, err := q.es.Search().
		Index(elasticsearch.ElasticIndexName).
		Query(esQuery).
		Do(ctx)

	if err != nil{
		return nil, err
	}

	return result, nil


}
