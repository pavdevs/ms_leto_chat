package chatsdto

type GetChatsResponse struct {
	Chats []ChatDTO `json:"items"`
}

func NewGetChatsResponse(chats []ChatDTO) *GetChatsResponse {

	return &GetChatsResponse{
		Chats: chats,
	}
}
