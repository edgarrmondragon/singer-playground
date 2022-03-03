package singer

import (
	"testing"
)

func TestStringType(t *testing.T) {
	type_object := jsonSchemaType{
		Type: "string",
	}

	if type_object.Type != "string" {
		t.Errorf("Expected %s, got %s", "string", type_object.Type)
	}
}

func TestDatetimeType(t *testing.T) {
	type_object := jsonSchemaType{
		Type:   "string",
		Format: "date-time",
	}

	if type_object.Type != "string" {
		t.Errorf("Expected %s, got %s", "string", type_object.Type)
	}

	if type_object.Format != "date-time" {
		t.Errorf("Expected %s, got %s", "date-time", type_object.Type)
	}
}

func TestObjectType(t *testing.T) {
	type_object := jsonSchemaType{
		Type: "object",
		Properties: map[string]jsonSchemaType{
			"id": {
				Type: "integer",
			},
		},
	}

	if type_object.Type != "object" {
		t.Errorf("Expected %s, got %s", "object", type_object.Type)
	}

	if type_object.Properties["id"].Type != "integer" {
		t.Errorf("Expected %s, got %s", "integer", type_object.Properties["id"].Type)
	}
}

func TestArrayType(t *testing.T) {
	type_object := jsonSchemaType{
		Type: "array",
		Items: &jsonSchemaType{
			Type: "integer",
		},
	}

	if type_object.Type != "array" {
		t.Errorf("Expected %s, got %s", "array", type_object.Type)
	}

	if type_object.Items.Type != "integer" {
		t.Errorf("Expected %s, got %s", "integer", type_object.Items.Type)
	}
}

func TestReadFromFile(t *testing.T) {
	schema := ReadSchemaFromFile("testdata/example_schema.json")

	if schema.Properties["id"].Type != "integer" {
		t.Errorf("Expected %s, got %s", "integer", schema.Properties["id"].Type)
	}

	if schema.Properties["name"].Type != "string" {
		t.Errorf("Expected %s, got %s", "string", schema.Properties["name"].Type)
	}

	if schema.Properties["created_at"].Type != "string" {
		t.Errorf("Expected %s, got %s", "string", schema.Properties["created_at"].Type)
	}

	if schema.Properties["created_at"].Format != "date-time" {
		t.Errorf("Expected %s, got %s", "date-time", schema.Properties["name"].Format)
	}
}
