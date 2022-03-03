package singer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type jsonSchemaType struct {
	Type       interface{}               `json:"type"`
	Format     string                    `json:"format"`
	Items      *jsonSchemaType           `json:"items,omitempty"`
	Properties map[string]jsonSchemaType `json:"properties,omitempty"`
}

type Schema struct {
	Properties map[string]jsonSchemaType `json:"properties"`
}

func ReadSchemaFromFile(filename string) Schema {
	var contents Schema

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}

	err = json.Unmarshal(data, &contents)
	if err != nil {
		fmt.Println(err)
	}
	return contents
}
