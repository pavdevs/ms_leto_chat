package messages

type Message struct {
	ID        int    `json:"id"`
	ChatID    int    `json:"channel_id"`
	UserID    int    `json:"user_id"`
	Text      string `json:"text"`
	CreatedAt string `json:"created_at"`
}
