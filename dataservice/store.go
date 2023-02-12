package dataservice

import "gorm.io/gorm"

type Store interface{
	Querier
}

type PGStore struct{
	Db *gorm.DB
	*Queries
}

func NewStore(db *gorm.DB) Store{
	return &PGStore{
		Db: db,
		Queries: NewQueries(db),
	}
}
