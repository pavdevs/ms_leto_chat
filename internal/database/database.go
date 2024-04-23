package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DBService struct {
	db     *sql.DB
	config Config
	logger *logrus.Logger
}

func NewDB(config Config, logger *logrus.Logger) *DBService {
	return &DBService{
		config: config,
		logger: logger,
	}
}

func (dbs *DBService) Connect() error {
	cs := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", dbs.config.host, dbs.config.user, dbs.config.dbName, "disable", dbs.config.password)
	database, err := sql.Open("postgres", cs)

	if err != nil {
		dbs.logger.Error(err)
		return err
	}

	err = database.Ping()

	if err != nil {
		dbs.logger.Error(err)
		return err
	}

	dbs.db = database

	return nil
}

func (dbs *DBService) Disconnect() error {
	return dbs.db.Close()
}
