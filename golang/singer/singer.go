package singer

import (
	"encoding/json"
	"time"
)

const (
	SchemaMessage string = "SCHEMA"
	RecordMessage string = "RECORD"
)

type SingerRecord struct {
	Stream        string          `json:"stream"`
	Record        json.RawMessage `json:"record"`
	TimeExtracted *time.Time      `json:"time_extracted"`
	Type          string          `json:"type"`
}

func NewSingerRecord(streamName string, record json.RawMessage) *SingerRecord {
	return &SingerRecord{
		Stream:        streamName,
		Record:        record,
		TimeExtracted: nil,
		Type:          RecordMessage,
	}
}

type SingerSchema struct {
	Stream             string          `json:"stream"`
	Schema             json.RawMessage `json:"schema"`
	Type               string          `json:"type"`
	KeyProperties      []string        `json:"key_properties"`
	BookmarkProperties []string        `json:"bookmark_properties"`
}

func NewSingerSchema(streamName string, schema json.RawMessage) *SingerSchema {
	return &SingerSchema{
		Stream:        streamName,
		Schema:        schema,
		Type:          SchemaMessage,
		KeyProperties: make([]string, 0),
	}
}
