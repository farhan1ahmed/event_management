package dataservice

import (
	"database/sql"
	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

type DatabaseMethods interface {
	AutoMigrate(dst ...interface{}) error
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Model(value interface{}) (tx *gorm.DB)
	Get(key string) (interface{}, bool)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}
type ElasticSearchMethods interface {
	Index() *elastic.IndexService
	Search(indices ...string) *elastic.SearchService
}

type Queries struct{
	db DatabaseMethods
	es ElasticSearchMethods
}

func NewQueries(db DatabaseMethods, es ElasticSearchMethods) *Queries{
	return &Queries{db:db, es:es}
}