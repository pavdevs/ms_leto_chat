package eventsparser

import "encoding/json"

type Event struct {
	Name    string          `json:"name"`
	Payload json.RawMessage `json:"payload"`
}
