package chats

import (
	"MsLetoChat/internal/services/chats"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ChatsAPI struct {
	logger *logrus.Logger
	cs     *chats.ChatsService
}

func NewChatsAPI(logger *logrus.Logger, cs *chats.ChatsService) *ChatsAPI {

	return &ChatsAPI{
		logger: logger,
		cs:     cs,
	}
}

func (c *ChatsAPI) CreateChat(w http.ResponseWriter, r *http.Request) {

}

func (c *ChatsAPI) DeleteChat(w http.ResponseWriter, r *http.Request) {

}

func (c *ChatsAPI) EditChat(w http.ResponseWriter, r *http.Request) {

}

func (c *ChatsAPI) GetChat(w http.ResponseWriter, r *http.Request) {

}

func (c *ChatsAPI) GetChats(w http.ResponseWriter, r *http.Request) {

}
