package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type envFile struct {
	DbName     string
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     string
	DbPoolSize int
	BuildEnv   string
	ServerHost string
	ServerPort string
	ElasticURL string
}

func (env *envFile) GetServerAddress() string {
	return env.ServerHost + ":" + env.ServerPort
}

func (env *envFile) GetDBUrl() string{
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		env.DbUsername, env.DbPassword, env.DbHost, env.DbPort, env.DbName)
	return dsn
}

var Env *envFile

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	dbPoolSize, err := strconv.Atoi(os.Getenv("DB_POOL_SIZE"))
	if err != nil {
		log.Println("Error parsing dbPoolSize")
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "event_ticket_service"
	}
	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		dbUsername = "postgres"
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"

	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}
	serverHost := os.Getenv("SERVER_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"

	}
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8081"
	}
	buildEnv := os.Getenv("BUILD_ENV")
	if buildEnv == "" {
		buildEnv = "dev"
	}
	elasticURL := os.Getenv("ELASTIC_URL")
	if elasticURL ==""{
		elasticURL = "http://0.0.0.0:9200"
	}

	Env = &envFile{
		DbName:     dbName,
		DbUsername: dbUsername,
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbPoolSize: dbPoolSize,
		BuildEnv:   buildEnv,
		ServerHost: serverHost,
		ServerPort: serverPort,
		ElasticURL: elasticURL,
	}
}