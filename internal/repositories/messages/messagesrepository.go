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

func (mr *MessagesRepository) AddMessage(userId int, message string) error {

	return nil
}

func (mr *MessagesRepository) DeleteMessage(userId int, messageId int) error {

	return nil
}

func (mr *MessagesRepository) GetMessages(userId int) ([]string, error) {

	return nil, nil
}
