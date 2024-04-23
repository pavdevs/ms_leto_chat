package repositories

import (
	"MsLetoChat/internal/database"
	"MsLetoChat/internal/repositories/chats"
	"MsLetoChat/internal/repositories/messages"
	"github.com/sirupsen/logrus"
)

type RepositoriesManager struct {
	cr *chats.ChatsRepository
	mr *messages.MessagesRepository
}

func NewRepositoriesManager(db *database.DBService, logger *logrus.Logger) *RepositoriesManager {

	return &RepositoriesManager{
		cr: chats.NewChatsRepository(db, logger),
		mr: messages.NewMessagesRepository(db, logger),
	}
}
