package messages

import "MsLetoChat/internal/repositories"

type MessagesService struct {
	rpm *repositories.RepositoriesManager
}

func NewMessagesService(rpm *repositories.RepositoriesManager) *MessagesService {

	return &MessagesService{
		rpm: rpm,
	}
}
