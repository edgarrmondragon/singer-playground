package singer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	SchemaMessage string = "SCHEMA"
	RecordMessage string = "RECORD"
)

type singerMessage struct {
	Stream string `json:"stream"`
	Type   string `json:"type"`
}

type singerSchema struct {
	Schema             json.RawMessage `json:"schema"`
	KeyProperties      []string        `json:"key_properties"`
	BookmarkProperties []string        `json:"bookmark_properties"`
	singerMessage
}

type singerRecord struct {
	singerMessage
	Record        json.RawMessage `json:"record"`
	TimeExtracted *time.Time      `json:"time_extracted"`
}

func PrintMessage(message interface{}) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, string(data))
}

func NewSingerRecord(streamName string, record json.RawMessage) *singerRecord {
	return &singerRecord{
		Record:        record,
		TimeExtracted: nil,
		singerMessage: singerMessage{
			Type:   RecordMessage,
			Stream: streamName,
		},
	}
}

func NewSingerSchema(streamName string, schema json.RawMessage) *singerSchema {
	return &singerSchema{
		Schema:        schema,
		KeyProperties: make([]string, 0),
		singerMessage: singerMessage{
			Type:   SchemaMessage,
			Stream: streamName,
		},
	}
}
