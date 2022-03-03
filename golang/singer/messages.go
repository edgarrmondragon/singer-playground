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

type Record map[string]interface{}

type singerMessage struct {
	Stream string `json:"stream"`
	Type   string `json:"type"`
}

type singerSchema struct {
	Schema             Schema   `json:"schema"`
	KeyProperties      []string `json:"key_properties"`
	BookmarkProperties []string `json:"bookmark_properties"`
	singerMessage
}

type singerRecord struct {
	singerMessage
	Record        Record     `json:"record"`
	TimeExtracted *time.Time `json:"time_extracted"`
}

func PrintMessage(message interface{}) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stdout, string(data))
}

func NewSingerRecord(streamName string, record Record) *singerRecord {
	return &singerRecord{
		Record:        record,
		TimeExtracted: nil,
		singerMessage: singerMessage{
			Type:   RecordMessage,
			Stream: streamName,
		},
	}
}

func NewSingerSchema(streamName string, schema Schema) *singerSchema {
	return &singerSchema{
		Schema:        schema,
		KeyProperties: make([]string, 0),
		singerMessage: singerMessage{
			Type:   SchemaMessage,
			Stream: streamName,
		},
	}
}
