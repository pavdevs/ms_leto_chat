package chatsdto

type GetChatRequest struct {
	ChatID int `json:"chat_id"`
}

type GetChatResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	OwnerID   int64  `json:"owner_id"`
	CreatedAt int64  `json:"created_at"`
}
