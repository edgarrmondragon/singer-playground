package singer

import (
	"testing"
)

func TestCatalog(t *testing.T) {
	catalog := SingerCatalog{
		Streams: []CatalogEntry{
			{
				Name: "jobs",
				Schema: Schema{
					Properties: map[string]jsonSchemaType{
						"id": {
							Type:   "string",
							Format: "uuid",
						},
						"name": {
							Type: "string",
						},
					},
				},
				StreamMetadata: []MetadataEntry{
					{
						Breadcrumb: []string{},
						Metadata: Metadata{
							Inclusion:         "available",
							ReplicationMethod: "FULL_TABLE",
						},
					},
					{
						Breadcrumb: []string{"properties", "id"},
						Metadata: Metadata{
							Inclusion: "automatic",
						},
					},
					{
						Breadcrumb: []string{"properties", "name"},
						Metadata: Metadata{
							Inclusion: "available",
						},
					},
				},
			},
		},
	}

	if catalog.Streams[0].Name != "jobs" {
		t.Errorf("Expected %s, got %s", "jobs", catalog.Streams[0].Name)
	}

	foundStream := catalog.FindStream("jobs")
	if foundStream.Name != "jobs" {
		t.Errorf("Expected to find stream %s", "jobs")
	}

	missingStream := catalog.FindStream("missing")
	if missingStream.Name != "" {
		t.Errorf("Expected not to find stream %s", "missing")
	}

	if foundStream.FindMetadata([]string{}).Metadata.ReplicationMethod != "FULL_TABLE" {
		t.Errorf("Expected not to find %s", "FULL_TABLE")
	}

	if foundStream.FindMetadata([]string{"properties", "id"}).Metadata.Inclusion != "automatic" {
		t.Errorf("Expected not to find %s", "automatic")
	}
}
