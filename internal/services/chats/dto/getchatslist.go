package chatsservicedto

type Chat struct {
	ID        int64
	Title     string
	OwnerID   int64
	CreatedAt int64
}

type GetChatsListServiceResponseDTO struct {
	Chats []Chat
}

func NewGetChatsListServiceResponseDTO(chats []Chat) *GetChatsListServiceResponseDTO {

	return &GetChatsListServiceResponseDTO{
		Chats: chats,
	}
}
