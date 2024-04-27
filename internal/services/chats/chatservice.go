package chats

import "MsLetoChat/internal/repositories"

type ChatsService struct {
	rpm *repositories.RepositoriesManager
}

func NewChatsService(rpm *repositories.RepositoriesManager) *ChatsService {

	return &ChatsService{
		rpm: rpm,
	}
}
