package repositories

import (
	"MsLetoChat/internal/database"
	"MsLetoChat/internal/repositories/chats"
	"MsLetoChat/internal/repositories/messages"
	"github.com/sirupsen/logrus"
)

type RepositoriesManager struct {
	Cr *chats.ChatsRepository
	Mr *messages.MessagesRepository
}

func NewRepositoriesManager(db *database.DBService, logger *logrus.Logger) *RepositoriesManager {

	return &RepositoriesManager{
		Cr: chats.NewChatsRepository(db, logger),
		Mr: messages.NewMessagesRepository(db, logger),
	}
}
