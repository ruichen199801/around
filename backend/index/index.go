package main

import (
        "context"
        "fmt"
        
        "github.com/olivere/elastic/v7"
)

const (
	POST_INDEX = "post"
	USER_INDEX = "user"
	ES_URL     = "http://10.128.0.2:9200" 
)

func main() {
    client, err := elastic.NewClient(
        elastic.SetURL(ES_URL),
        elastic.SetBasicAuth("elastic", "placeholder")) 
    if err != nil {
        panic(err)
    }

    exists, err := client.IndexExists(POST_INDEX).Do(context.Background())
    if err != nil {
        panic(err)
    }

    if !exists {
        mapping := `{
            "mappings": {
                "properties": {
                    "id":       { "type": "keyword" },
                    "user":     { "type": "keyword" },
                    "message":  { "type": "text" },
                    "url":      { "type": "keyword", "index": false },
                    "type":     { "type": "keyword", "index": false }
                }
            }
        }`
        _, err := client.CreateIndex(POST_INDEX).Body(mapping).Do(context.Background())
        if err != nil {
            panic(err)
        }
    }

    exists, err = client.IndexExists(USER_INDEX).Do(context.Background())
    if err != nil {
        panic(err)
    }

    if !exists {
        mapping := `{
                        "mappings": {
                                "properties": {
                                        "username": {"type": "keyword"},
                                        "password": {"type": "keyword"},
                                        "age":      {"type": "long", "index": false},
                                        "gender":   {"type": "keyword", "index": false}
                                }
                        }
                }`
        _, err = client.CreateIndex(USER_INDEX).Body(mapping).Do(context.Background())
        if err != nil {
            panic(err)
        }
    }
    fmt.Println("Indexes are created.")
}
