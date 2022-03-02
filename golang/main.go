package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"tap-data-jobs/singer"
	"time"
)

const (
	baseURL string = "https://api.datastackjobs.com"
	version string = "0.0.1"
)

type meta struct {
	After string `json:"after"`
}

type page struct {
	Data []json.RawMessage `json:"data"`
	Meta meta              `json:"meta"`
}

type singerStream struct {
	Name   string
	Path   string
	Schema json.RawMessage
}

var streams []singerStream

func readJSONFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	return data
}

func (stream *singerStream) getCatalogEntry() singer.CatalogEntry {
	metadata := make([]singer.MetadataEntry, 0)
	metadata = append(
		metadata,
		singer.MetadataEntry{
			Breadcrumb: []string{},
			Metadata: singer.Metadata{
				Selected: true,
			},
		},
	)
	return singer.CatalogEntry{
		Name:           stream.Name,
		Schema:         stream.Schema,
		StreamMetadata: metadata,
	}
}

func syncRecords() {
	var resp *http.Response
	client := &http.Client{}

	for _, stream := range streams {
		url := baseURL + stream.Path
		for {
			fmt.Fprintf(os.Stderr, "Requesting data from %s\n", url)
			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				fmt.Println("error:", err)
			}

			req.Header.Add("User-Agent", fmt.Sprintf("tap-datastack-jobs/%s", version))

			extracted := time.Now()
			resp, err = client.Do(req)
			if err != nil {
				fmt.Println("error:", err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("error:", err)
			}

			page := page{}
			err = json.Unmarshal(body, &page)
			if err != nil {
				log.Fatal(err)
			}

			for _, object := range page.Data {
				recordMessage := singer.NewSingerRecord(stream.Name, object)
				recordMessage.TimeExtracted = &extracted
				singer.PrintMessage(recordMessage)
			}

			if page.Meta.After != "" {
				url = baseURL + "/v1/jobs?page[after]=" + page.Meta.After
			} else {
				break
			}
		}
	}

}

func getDefaultCatalog(selectStreams bool) singer.SingerCatalog {
	entries := make([]singer.CatalogEntry, 0)

	for _, stream := range streams {
		entry := stream.getCatalogEntry()
		entry.StreamMetadata[0].Metadata.Selected = selectStreams

		entries = append(entries, entry)
	}

	return singer.SingerCatalog{
		Streams: entries,
	}
}

func doDiscover() {
	catalog := getDefaultCatalog(false)
	data, err := json.Marshal(catalog)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stdout, string(data))
	os.Exit(0)
}

func doSync(catalog singer.SingerCatalog) {
	for _, stream := range streams {
		entry := catalog.FindStream(stream.Name)
		if entry.Name != "" {
			stream_metadata := entry.FindMetadata([]string{})
			if stream_metadata.Metadata.Selected {
				fmt.Fprintf(os.Stderr, "Syncing '%s'\n", stream.Name)
				schemaMessage := singer.NewSingerSchema(stream.Name, entry.Schema)
				schemaMessage.KeyProperties = []string{"id"}
				singer.PrintMessage(schemaMessage)
				syncRecords()
			} else {
				fmt.Fprintf(os.Stderr, "Stream '%s' is not selected in catalog\n", stream.Name)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Stream '%s' is not present in catalog\n", stream.Name)
		}
	}
}

func main() {
	var catalogFile string
	var catalog singer.SingerCatalog
	var discover bool

	streams = append(streams, singerStream{
		Name:   "jobs",
		Path:   "/v1/jobs",
		Schema: readJSONFile("./job.json"),
	})

	flag.StringVar(&catalogFile, "p", "", "Specify catalog. Default is root")
	flag.BoolVar(&discover, "d", false, "Run in discovery mode.")
	flag.Parse()

	if discover {
		doDiscover()
	}

	if catalogFile != "" {
		catalog = singer.ReadCatalogFromFile(catalogFile)
	} else {
		catalog = getDefaultCatalog(true)

	}

	doSync(catalog)
}
