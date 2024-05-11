package chats

import (
	"MsLetoChat/internal/database"
	chatrepositorydto "MsLetoChat/internal/repositories/chats/dto"
	"database/sql"
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

func (mr *ChatsRepository) CreateChat(chat chatrepositorydto.CreateChatDTO) (*chatrepositorydto.ChatResponseDTO, error) {
	db, err := mr.getDB()

	// Проверка соединения с базой данных
	if err != nil {
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
	err = db.QueryRow(q, chat.Title, chat.OwnerID, createdAt).Scan(&response.ChatID, &response.Title, &response.OwnerID, &response.CreatedAt)
	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to create chat: %w", err)
	}

	// Возврат полной информации о созданном чате
	return &response, nil
}

func (mr *ChatsRepository) GetChat(chatID int64) (*chatrepositorydto.ChatDTO, error) {
	db, err := mr.getDB()

	// Проверка соединения с базой данных
	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	var chat chatrepositorydto.ChatDTO
	q := `SELECT id, title, owner_id, created_at FROM chats WHERE id = $1`
	err = db.QueryRow(q, chatID).Scan(&chat.ChatID, &chat.Title, &chat.OwnerID, &chat.CreatedAt)

	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}

	return &chat, nil
}

func (mr *ChatsRepository) UpdateChat(title string, id int64) (*chatrepositorydto.ChatDTO, error) {
	db, err := mr.getDB()

	// Проверка соединения с базой данных
	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	var chat chatrepositorydto.ChatDTO
	q := `UPDATE chats SET title = $1 WHERE id = $2 RETURNING id, title, owner_id, created_at`
	err = db.QueryRow(q, title, id).Scan(&chat.ChatID, &chat.Title, &chat.OwnerID, &chat.CreatedAt)

	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to update chat: %w", err)
	}

	return &chat, nil
}

func (mr *ChatsRepository) DeleteChat(chatID int64) error {
	db, err := mr.getDB()

	// Проверка соединения с базой данных
	if err != nil {
		mr.logger.Error(err)
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	q := `DELETE FROM chats WHERE id = $1`
	_, err = db.Exec(q, chatID)

	if err != nil {
		mr.logger.Error(err)
		return fmt.Errorf("failed to delete chat: %w", err)
	}

	return nil
}

func (mr *ChatsRepository) GetChatsList(ownerID int64) (*chatrepositorydto.GetChatsListRepositoryResponseDTO, error) {
	db, err := mr.getDB()

	// Проверка соединения с базой данных
	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	q := `SELECT id, title, owner_id, created_at FROM chats WHERE owner_id = $1`
	var chatsRows []chatrepositorydto.ChatDTO

	rows, err := db.Query(q, ownerID)

	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to get chats list: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var chat chatrepositorydto.ChatDTO
		err := rows.Scan(&chat.ChatID, &chat.Title, &chat.OwnerID, &chat.CreatedAt)
		if err != nil {
			mr.logger.Error(err)
			return nil, fmt.Errorf("failed to get chats list: %w", err)
		}

		chatsRows = append(chatsRows, chat)
	}

	err = rows.Err()

	if err != nil {
		mr.logger.Error(err)
		return nil, fmt.Errorf("failed to get chats list: %w", err)
	}

	return chatrepositorydto.NewGetChatsListRepositoryResponseDTO(chatsRows), nil
}

func (mr *ChatsRepository) getDB() (*sql.DB, error) {
	db := mr.db.GetDB()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
