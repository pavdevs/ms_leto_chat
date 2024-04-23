package main

import (
	"MsLetoChat/internal/database"
	"MsLetoChat/internal/repositories"
	"MsLetoChat/internal/server"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
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
	wsServer := prepareServer(logger)

	err := dbService.Connect()

	if err != nil {
		logger.Fatal(err)
	}

	http.Handle("/ws", websocket.Handler(wsServer.HandleWS))
	logger.Info(fmt.Sprintf("Server listening on https://%s:%s", os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")))

	httpCs := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))

	if err := http.ListenAndServe(httpCs, nil); err != nil {
		logger.Error(err)
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

func prepareServer(logger *logrus.Logger) *server.WSServer {

	return server.NewServer(
		server.NewConfig(
			os.Getenv("SERVER_HOST"),
			os.Getenv("SERVER_PORT"),
		),
		logger,
		[]byte(os.Getenv("SECRET_KEY")),
	)
}
