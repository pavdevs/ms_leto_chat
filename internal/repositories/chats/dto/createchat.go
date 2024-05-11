package chatrepositorydto

type CreateChatDTO struct {
	Title   string
	OwnerID int64
}

func NewCreateChatDTO(title string, ownerID int64) CreateChatDTO {
	return CreateChatDTO{
		Title:   title,
		OwnerID: ownerID,
	}
}

type ChatResponseDTO struct {
	ChatID    int64
	Title     string
	OwnerID   int64
	CreatedAt int64
}

func NewChatResponseDTO(chatID int64, title string, ownerID int64, createdAt int64) *ChatResponseDTO {

	return &ChatResponseDTO{
		ChatID:    chatID,
		Title:     title,
		OwnerID:   ownerID,
		CreatedAt: createdAt,
	}
}
