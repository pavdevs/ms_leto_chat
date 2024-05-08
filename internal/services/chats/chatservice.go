package chats

import (
	"MsLetoChat/internal/repositories"
	chatrepositorydto "MsLetoChat/internal/repositories/chats/dto"
	chatsservicedto "MsLetoChat/internal/services/chats/dto"
	"fmt"
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

func (s *ChatsService) CreateChat(chat chatsservicedto.ChatDTO) (*chatsservicedto.ChatResponseDTO, error) {

	chatRepDTO := chatrepositorydto.NewChatDTO(chat.Title, chat.OwnerID)

	c, err := s.rpm.Cr.CreateChat(chatRepDTO)

	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return chatsservicedto.NewChatResponseDTO(
			c.ChatID,
			c.Title,
			c.OwnerID,
			c.CreatedAt,
		),
		nil
}

func (s *ChatsService) DeleteChat(chatID int64) error {
	return s.rpm.Cr.DeleteChat(chatID)
}

func (s *ChatsService) GetChatsList(ownerID int64) (*chatsservicedto.GetChatsListServiceResponseDTO, error) {

	c, err := s.rpm.Cr.GetChatsList(ownerID)

	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	var items []chatsservicedto.Chat

	for item := range c.Chats {
		items = append(items, chatsservicedto.Chat{
			ID:        c.Chats[item].ID,
			Title:     c.Chats[item].Title,
			OwnerID:   c.Chats[item].OwnerID,
			CreatedAt: c.Chats[item].CreatedAt,
		})
	}

	return chatsservicedto.NewGetChatsListServiceResponseDTO(items), nil
}
