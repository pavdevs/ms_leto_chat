package chatsdto

type DeleteChatRequest struct {
	ChatID int `json:"chat_id"`
}

type DeleteChatResponse struct {
	Message string `json:"message"`
}
