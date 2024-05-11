package chatsservicedto

type CreateChatDTO struct {
	Title   string
	OwnerID int64
}

func NewCreateChatDTO(title string, ownerID int64) ChatDTO {

	return ChatDTO{
		Title:   title,
		OwnerID: ownerID,
	}
}
