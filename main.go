package main

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	people := []Person{
		{Name: "John Doe", Age: 30, Gender: "Male"},
		{Name: "Jane Doe", Age: 28, Gender: "Female"},
	}

	for _, person := range people {
		personJson, err := json.Marshal(person)
		if err != nil {
			log.Fatalf("Error marshaling the person: %s", err)
		}

		req := esapi.IndexRequest{
			Index:      "test-index",
			DocumentID: person.Name,
			Body:       strings.NewReader(string(personJson)),
		}

		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Println(res.String())
		} else {
			log.Printf("Successfully indexed document: %s", person.Name)
		}
	}
}
