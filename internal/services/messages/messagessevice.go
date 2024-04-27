package messages

import (
	"MsLetoChat/internal/repositories"
	"github.com/sirupsen/logrus"
)

type MessagesService struct {
	logger *logrus.Logger
	rpm    *repositories.RepositoriesManager
}

func NewMessagesService(logger *logrus.Logger, rpm *repositories.RepositoriesManager) *MessagesService {

	return &MessagesService{
		logger: logger,
		rpm:    rpm,
	}
}
