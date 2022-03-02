package singer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Metadata struct {
	Inclusion          string   `json:"inclusion"`
	Selected           bool     `json:"selected"`
	ReplicationMethod  string   `json:"replication-method"`
	TableKeyProperties []string `json:"table-key-properties"`
}

type MetadataEntry struct {
	Breadcrumb []string `json:"breadcrumb"`
	Metadata   Metadata `json:"metadata"`
}

type CatalogEntry struct {
	Name           string          `json:"name"`
	Schema         json.RawMessage `json:"schema"`
	StreamMetadata []MetadataEntry `json:"metadata"`
}

type SingerCatalog struct {
	Streams []CatalogEntry `json:"streams"`
}

func ReadCatalogFromFile(filename string) SingerCatalog {
	var catalog SingerCatalog

	catalogData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}

	err = json.Unmarshal(catalogData, &catalog)
	if err != nil {
		fmt.Println(err)
	}

	return catalog
}

func (catalog SingerCatalog) FindStream(name string) CatalogEntry {
	for _, stream := range catalog.Streams {
		if stream.Name == name {
			return stream
		}
	}
	return CatalogEntry{}
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func (entry CatalogEntry) FindMetadata(breadcrumb []string) MetadataEntry {
	for _, metadata := range entry.StreamMetadata {
		if stringSlicesEqual(metadata.Breadcrumb, breadcrumb) {
			return metadata
		}
	}
	return MetadataEntry{}
}
