package eventsparser

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

type EventsParser struct {
	logger *logrus.Logger
}

func NewEventsParser(logger *logrus.Logger) *EventsParser {
	return &EventsParser{
		logger: logger,
	}
}

// ParseEvent анализирует сообщение и обрабатывает payload в зависимости от name
func (p *EventsParser) ParseEvent(message []byte) error {
	var event Event
	if err := json.Unmarshal(message, &event); err != nil {
		p.logger.Error(err)
		return fmt.Errorf("failed to unmarshal event: %w", err)
	}

	switch event.Name {
	case "user_login":
		return nil
	case "user_logout":
		//var payload UserLogout
		//if err := json.Unmarshal(event.Payload, &payload); err != nil {
		//	p.logger.Error(err)
		//	return fmt.Errorf("failed to unmarshal event: %w", err)
		//}
		// обработка логики user_logout
		//fmt.Printf("User logged out: %+v\n", payload)
		return nil
	case "message":
		return nil
	default:
		fmt.Println("Unknown event")
	}

	return nil
}
