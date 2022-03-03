package singer

import (
	"testing"
)

func TestNewSingerRecord(t *testing.T) {
	streamName := "test"
	record := make(map[string]interface{})
	record["id"] = 1
	record["name"] = "Bob"

	message := NewSingerRecord(streamName, record)

	if message.Type != RecordMessage {
		t.Errorf("Expected %s, got %s", RecordMessage, message.Type)
	}

	if message.Stream != streamName {
		t.Errorf("Expected %s, got %s", streamName, message.Stream)
	}

	if message.Record["id"] != record["id"] {
		t.Errorf("Expected %s, got %s", record["id"], message.Record["id"])
	}

	if message.Record["name"] != record["name"] {
		t.Errorf("Expected %s, got %s", record["name"], message.Record["name"])
	}
}

func TestNewSingerSchema(t *testing.T) {
	streamName := "test"
	schema := Schema{
		Properties: map[string]jsonSchemaType{
			"id": {
				Type: "integer",
			},
			"name": {
				Type: []interface{}{"string", "null"},
			},
		},
	}

	message := NewSingerSchema(streamName, schema)

	if message.Type != SchemaMessage {
		t.Errorf("Expected %s, got %s", SchemaMessage, message.Type)
	}

	if message.Stream != streamName {
		t.Errorf("Expected %s, got %s", streamName, message.Stream)
	}

	propId := schema.Properties["id"]
	switch dt := propId.Type.(type) {
	case string:
		if dt != "integer" {
			t.Errorf("Expected %s, got %s", "integer", dt)
		}
	default:
		t.Errorf("%s is not of the right type", dt)
	}

	propName := schema.Properties["name"]
	switch dt := propName.Type.(type) {
	case []interface{}:
		if dt[0] != "string" && dt[1] != "null" {
			t.Errorf("Expected %s, got %s", []string{"string", "null"}, dt)
		}
	default:
		t.Errorf("%T is not of the right type", dt)
	}
}
