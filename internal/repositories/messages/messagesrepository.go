package messages

import (
	"MsLetoChat/internal/database"
	"github.com/sirupsen/logrus"
)

type MessagesRepository struct {
	db     *database.DBService
	logger *logrus.Logger
}

func NewMessagesRepository(db *database.DBService, logger *logrus.Logger) *MessagesRepository {

	return &MessagesRepository{
		db:     db,
		logger: logger,
	}
}
