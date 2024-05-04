package chatsdto

type CreateChatRequest struct {
	Title   string `json:"title"`
	OwnerID int64  `json:"owner_id"`
}

type CreateChatResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	OwnerID   int64  `json:"owner_id"`
	CreatedAt int64  `json:"created_at"`
}

func NewCreateChatResponse(id int64, title string, ownerID int64, createdAt int64) *CreateChatResponse {

	return &CreateChatResponse{
		ID:        id,
		Title:     title,
		OwnerID:   ownerID,
		CreatedAt: createdAt,
	}
}
