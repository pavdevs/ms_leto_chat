package main

import (
	"MsLetoChat/internal/api/chats"
	"MsLetoChat/internal/api/messages"
	"MsLetoChat/internal/database"
	"MsLetoChat/internal/repositories"
	"MsLetoChat/internal/server"
	chatsservice "MsLetoChat/internal/services/chats"
	messagesservice "MsLetoChat/internal/services/messages"
	"MsLetoChat/internal/support/eventsparser"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = prepareLogger()
	loadEnv(logger)
}

// @title Leto chats service
// @version 1.0
// @description API Server for leto app for chats api

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	dbService := prepareDatabase(logger)
	rpm := prepareRepositories(dbService, logger)

	chatsApi := prepareChatsModule(logger, rpm)
	messagesApi := prepareMessagesModule(logger, rpm)
	eventsParser := prepareEventsParser(logger)

	err := dbService.Connect()

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Database connected")

	config := server.NewConfig("", "8080")
	wsServer := server.NewServer(config, logger, chatsApi, messagesApi, eventsParser)

	logger.Infof("Server localhost started on port %s", config.Port)

	if err := wsServer.Start(); err != nil {
		logger.Fatal(err)
	}
}

func prepareLogger() *logrus.Logger {
	return logrus.New()
}

func loadEnv(logger *logrus.Logger) {
	if err := godotenv.Load("cmd/config.env"); err != nil {
		logger.Error("No .env file found")
	}
}

func prepareDatabase(logger *logrus.Logger) *database.DBService {
	dbConfig := database.NewConfig(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	dbService := database.NewDB(dbConfig, logger)

	return dbService
}

func prepareRepositories(db *database.DBService, logger *logrus.Logger) *repositories.RepositoriesManager {

	return repositories.NewRepositoriesManager(db, logger)
}

func prepareChatsModule(logger *logrus.Logger, rpm *repositories.RepositoriesManager) *chats.ChatsAPI {

	cs := chatsservice.NewChatsService(
		logger,
		rpm,
	)

	chatsApi := chats.NewChatsAPI(
		logger,
		cs,
	)

	return chatsApi
}

func prepareMessagesModule(logger *logrus.Logger, rpm *repositories.RepositoriesManager) *messages.MessagesApi {

	ms := messagesservice.NewMessagesService(
		logger,
		rpm,
	)

	messagesApi := messages.NewMessagesApi(
		logger,
		ms,
	)

	return messagesApi
}

func prepareEventsParser(logger *logrus.Logger) *eventsparser.EventsParser {

	return eventsparser.NewEventsParser(logger)
}
