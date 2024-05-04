package chatsdto

type GetChatsRequest struct {
	OwnerID int `json:"owner_id"`
}

type chat struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	OwnerID   int64  `json:"owner_id"`
	CreatedAt int64  `json:"created_at"`
}

type GetChatsResponse struct {
	Chats []chat `json:"chats"`
}
