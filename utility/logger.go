package utility

import (
	"event_ticket_service/env"
	"go.uber.org/zap"
)
const(
	DEVELOPMENT = "dev"
	PRODUCTION = "prod"
)

var logger *zap.Logger

func init(){
	switch env.Env.BuildEnv{
	case DEVELOPMENT:
		logger, _ = zap.NewDevelopment()
		return
	case PRODUCTION:
		logger, _ = zap.NewProduction()
		return
	default:
		logger = zap.NewExample()
		return
	}
}

func GetLogger() *zap.Logger{
	return logger
}