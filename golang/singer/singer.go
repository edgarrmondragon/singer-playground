package singer

import (
	"encoding/json"
)

const (
	SchemaMessage string = "SCHEMA"
	RecordMessage string = "RECORD"
)

type SingerRecord struct {
	Stream string          `json:"stream"`
	Record json.RawMessage `json:"record"`
	Type   string          `json:"type"`
}

func NewSingerRecord() *SingerRecord {
	return &SingerRecord{
		Stream: "",
		Record: make(json.RawMessage, 0),
		Type:   RecordMessage,
	}
}

type SingerSchema struct {
	Stream        string          `json:"stream"`
	Schema        json.RawMessage `json:"schema"`
	Type          string          `json:"type"`
	KeyProperties []string        `json:"key_properties"`
}

func NewSingerSchema() *SingerSchema {
	return &SingerSchema{
		Stream:        "",
		Schema:        make(json.RawMessage, 0),
		Type:          SchemaMessage,
		KeyProperties: make([]string, 0),
	}
}
