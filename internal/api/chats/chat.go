package chats

type Chat struct {
	ChatID    int    `json:"chat_id"`
	Title     string `json:"title"`
	OwnerID   string `json:"owner_id"`
	CreatedAt string `json:"created_at"`
}
