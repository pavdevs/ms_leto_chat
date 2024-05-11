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

func (s *ChatsService) CreateChat(chat chatsservicedto.ChatDTO) (*chatsservicedto.ChatDTO, error) {

	chatRepDTO := chatrepositorydto.NewCreateChatDTO(chat.Title, chat.OwnerID)

	c, err := s.rpm.Cr.CreateChat(chatRepDTO)

	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return &chatsservicedto.ChatDTO{
			ChatID:    c.ChatID,
			Title:     c.Title,
			OwnerID:   c.OwnerID,
			CreatedAt: c.CreatedAt,
		},
		nil
}

func (s *ChatsService) DeleteChat(chatID int64) error {
	return s.rpm.Cr.DeleteChat(chatID)
}

func (s *ChatsService) GetChat(chatID int64) (*chatsservicedto.ChatDTO, error) {

	c, err := s.rpm.Cr.GetChat(chatID)

	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	return &chatsservicedto.ChatDTO{
		ChatID:    c.ChatID,
		Title:     c.Title,
		OwnerID:   c.OwnerID,
		CreatedAt: c.CreatedAt,
	}, nil
}

func (s *ChatsService) UpdateChat(title string, id int64) (*chatsservicedto.ChatDTO, error) {
	chat, err := s.rpm.Cr.UpdateChat(title, id)

	if err != nil {
		return nil, fmt.Errorf("failed to update chat: %w", err)
	}

	return &chatsservicedto.ChatDTO{
		ChatID:    chat.ChatID,
		Title:     chat.Title,
		OwnerID:   chat.OwnerID,
		CreatedAt: chat.CreatedAt,
	}, nil
}

func (s *ChatsService) GetChatsList(ownerID int64) (*chatsservicedto.GetChatsListServiceResponseDTO, error) {

	c, err := s.rpm.Cr.GetChatsList(ownerID)

	if err != nil {
		s.logger.Error(err)
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	var items []chatsservicedto.ChatDTO

	for item := range c.Chats {
		items = append(items, chatsservicedto.ChatDTO{
			ChatID:    c.Chats[item].ChatID,
			Title:     c.Chats[item].Title,
			OwnerID:   c.Chats[item].OwnerID,
			CreatedAt: c.Chats[item].CreatedAt,
		})
	}

	return chatsservicedto.NewGetChatsListServiceResponseDTO(items), nil
}
