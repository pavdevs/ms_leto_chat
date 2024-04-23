package chats

import (
	"MsLetoChat/internal/database"
	"github.com/sirupsen/logrus"
)

type ChatsRepository struct {
	db     *database.DBService
	logger *logrus.Logger
}

func NewChatsRepository(db *database.DBService, logger *logrus.Logger) *ChatsRepository {

	return &ChatsRepository{
		db:     db,
		logger: logger,
	}
}
