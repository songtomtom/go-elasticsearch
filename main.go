package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Sentence struct {
	Content string `json:"content"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("You must provide text as a command line argument.")
	}
	text := os.Args[1]

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	sentences := strings.Split(text, ". ")
	for i, sentence := range sentences {
		s := Sentence{Content: sentence}
		sJson, err := json.Marshal(s)
		if err != nil {
			log.Fatalf("Error marshaling the sentence: %s", err)
		}

		req := esapi.IndexRequest{
			Index:      "test-index",
			DocumentID: strconv.Itoa(i),
			Body:       strings.NewReader(string(sJson)),
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Println(res.String())
		} else {
			log.Printf("Successfully indexed sentence: %s", sentence)
		}
	}
}
