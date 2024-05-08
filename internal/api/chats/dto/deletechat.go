package chatsdto

type DeleteChatResponse struct {
	Message string `json:"message"`
}

func NewDeleteChatResponse() *DeleteChatResponse {
	return &DeleteChatResponse{
		Message: "Chat deleted",
	}
}
