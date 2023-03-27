package dataservice

import (
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type Store interface{
	Querier
}

type PGStore struct{
	Db *gorm.DB
	Es *elastic.Client
	*Queries
}

func NewStore(db *gorm.DB, es *elastic.Client ) Store{
	return &PGStore{
		Db: db,
		Es: es,
		Queries: NewQueries(db, es),
	}
}
