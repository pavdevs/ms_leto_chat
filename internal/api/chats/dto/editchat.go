package chatsdto

type EditChatRequest struct {
	Title   string `json:"title"`
	OwnerID int64  `json:"owner_id"`
}

type EditChatResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	OwnerID   int64  `json:"owner_id"`
	CreatedAt int64  `json:"created_at"`
}
