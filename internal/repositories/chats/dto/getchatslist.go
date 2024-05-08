package chatrepositorydto

type Chat struct {
	ID        int64
	Title     string
	OwnerID   int64
	CreatedAt int64
}

type GetChatsListRepositoryResponseDTO struct {
	Chats []Chat
}

func NewGetChatsListRepositoryResponseDTO(chats []Chat) *GetChatsListRepositoryResponseDTO {

	return &GetChatsListRepositoryResponseDTO{
		Chats: chats,
	}
}
