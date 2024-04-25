package main

import (
	"MsLetoChat/internal/database"
	"MsLetoChat/internal/repositories"
	"MsLetoChat/internal/server"
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

func main() {
	dbService := prepareDatabase(logger)
	prepareRepositories(dbService, logger)

	err := dbService.Connect()

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Database connected")

	config := server.NewConfig("", "8080")
	wsServer := server.NewServer(config, logger)

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
