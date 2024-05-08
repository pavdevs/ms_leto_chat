package chatsdto

type Chat struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	OwnerID   int64  `json:"owner_id"`
	CreatedAt int64  `json:"created_at"`
}

type GetChatsResponse struct {
	Chats []Chat `json:"items"`
}

func NewGetChatsResponse(chats []Chat) *GetChatsResponse {

	return &GetChatsResponse{
		Chats: chats,
	}
}
