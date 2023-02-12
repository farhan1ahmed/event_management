package dataservice

import (
	"event_ticket_service/env"
	"event_ticket_service/utility"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresdb *gorm.DB

func init(){
	logger := utility.GetLogger()
	dsn := env.Env.GetDBUrl()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		logger.Error(err.Error())
	}
	postgresdb = db
}

func GetDBConnection() *gorm.DB{
	return postgresdb
}