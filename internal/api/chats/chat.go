package chats

type Chat struct {
	ChatID    int64  `json:"chat_id"`
	Title     string `json:"title"`
	OwnerID   string `json:"owner_id"`
	CreatedAt int64  `json:"created_at"`
}

func NewChat(chatID int64, title string, ownerID string, createdAt int64) Chat {

	return Chat{
		ChatID:    chatID,
		Title:     title,
		OwnerID:   ownerID,
		CreatedAt: createdAt,
	}
}
