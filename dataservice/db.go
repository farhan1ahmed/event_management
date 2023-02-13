package dataservice

import (
	"database/sql"
	"gorm.io/gorm"
)

type DatabaseMethods interface {
	AutoMigrate(dst ...interface{}) error
	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
	Model(value interface{}) (tx *gorm.DB)
	Get(key string) (interface{}, bool)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
}
type Queries struct{
	db DatabaseMethods
}

func NewQueries(db DatabaseMethods) *Queries{
	return &Queries{db:db}
}