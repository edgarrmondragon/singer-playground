package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tap-data-jobs/singer"
	"time"
)

const baseURL string = "https://api.datastackjobs.com"

type meta struct {
	After string `json:"after"`
}

type page struct {
	Data []json.RawMessage `json:"data"`
	Meta meta              `json:"meta"`
}

func main() {
	streamName := "jobs"
	schema, err := ioutil.ReadFile("./job.json")
	if err != nil {
		fmt.Print(err)
	}
	schemaMessage := singer.NewSingerSchema(streamName, schema)
	schemaMessage.KeyProperties = []string{"id"}

	data, err := json.Marshal(schemaMessage)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	var resp *http.Response
	client := &http.Client{}
	url := baseURL + "/v1/jobs"

	for {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("error:", err)
		}

		req.Header.Add("User-Agent", "tap-datastack-jobs/0.0.1")

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
			recordMessage := singer.NewSingerRecord(streamName, object)
			recordMessage.TimeExtracted = &extracted
			data, err = json.Marshal(recordMessage)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(data))
		}

		if page.Meta.After != "" {
			url = baseURL + "/v1/jobs?page[after]=" + page.Meta.After
		} else {
			break
		}
	}
}
