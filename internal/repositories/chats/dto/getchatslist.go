package chatrepositorydto

type GetChatsListRepositoryResponseDTO struct {
	Chats []ChatDTO
}

func NewGetChatsListRepositoryResponseDTO(chats []ChatDTO) *GetChatsListRepositoryResponseDTO {

	return &GetChatsListRepositoryResponseDTO{
		Chats: chats,
	}
}
