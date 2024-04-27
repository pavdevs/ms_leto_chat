package chats

import (
	"MsLetoChat/internal/repositories"
	"github.com/sirupsen/logrus"
)

type ChatsService struct {
	logger *logrus.Logger
	rpm    *repositories.RepositoriesManager
}

func NewChatsService(logger *logrus.Logger, rpm *repositories.RepositoriesManager) *ChatsService {

	return &ChatsService{
		logger: logger,
		rpm:    rpm,
	}
}
