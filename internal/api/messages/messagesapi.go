package messages

import (
	"MsLetoChat/internal/services/messages"
	"github.com/sirupsen/logrus"
)

type MessagesApi struct {
	logger *logrus.Logger
	ms     *messages.MessagesService
}

func NewMessagesApi(logger *logrus.Logger, ms *messages.MessagesService) *MessagesApi {

	return &MessagesApi{
		logger: logger,
		ms:     ms,
	}
}

func (ma *MessagesApi) AddMessage(msg *Message) error {

	return nil
}

func (ma *MessagesApi) EditMessage(text string) error {

	return nil
}

func (ma *MessagesApi) DeleteMessage(msgID int) error {

	return nil
}

func (ma *MessagesApi) GetMessages() ([]Message, error) {

	return nil, nil
}
