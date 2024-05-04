package chats

import (
	"MsLetoChat/internal/database"
	chatrepositorydto "MsLetoChat/internal/repositories/chats/dto"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type ChatsRepository struct {
	db     *database.DBService
	logger *logrus.Logger
}

func NewChatsRepository(db *database.DBService, logger *logrus.Logger) *ChatsRepository {

	return &ChatsRepository{
		db:     db,
		logger: logger,
	}
}

func (mr *ChatsRepository) CreateChat(chat chatrepositorydto.ChatDTO) (*chatrepositorydto.ChatResponseDTO, error) {
	db := mr.db.GetDB()

	// Проверка соединения с базой данных
	if err := db.Ping(); err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Формирование запроса на создание нового чата
	createdAt := time.Now().Unix()
	q := `INSERT INTO chats (title, owner_id, created_at) 
          VALUES ($1, $2, $3) 
          RETURNING id, title, owner_id, created_at`

	// Подготовка структуры для хранения результатов запроса
	var response chatrepositorydto.ChatResponseDTO
	err := db.QueryRow(q, chat.Title, chat.OwnerID, createdAt).Scan(&response.ChatID, &response.Title, &response.OwnerID, &response.CreatedAt)
	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	// Возврат полной информации о созданном чате
	return &response, nil
}

func (mr *ChatsRepository) DeleteChat(chatId int) error {

	return nil
}
