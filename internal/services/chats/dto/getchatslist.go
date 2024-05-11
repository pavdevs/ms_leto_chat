package chatsservicedto

type GetChatsListServiceResponseDTO struct {
	Chats []ChatDTO
}

func NewGetChatsListServiceResponseDTO(chats []ChatDTO) *GetChatsListServiceResponseDTO {

	return &GetChatsListServiceResponseDTO{
		Chats: chats,
	}
}
