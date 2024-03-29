package tools

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"log"

// 	"github.com/elastic/go-elasticsearch"

// 	//"strconv"
// 	"strings"
// 	//"github.com/elastic/go-elasticsearch/esapi"
// )

// func ESMain() {
// 	log.SetFlags(0)

// 	var (
// 		r map[string]interface{}
// 		//wg sync.WaitGroup
// 	)

// 	cfg := elasticsearch.Config{
// 		Addresses: []string{
// 			"http://192.168.1.211:8030",
// 			//"http://localhost:9201",
// 		},
// 	}
// 	// Initialize a client with the default settings.
// 	//
// 	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
// 	//
// 	//es, err := elasticsearch.NewDefaultClient()
// 	es, err := elasticsearch.NewClient(cfg)
// 	if err != nil {
// 		log.Fatalf("Error creating the client: %s", err)
// 	}

// 	// 1. Get cluster info
// 	//
// 	res, err := es.Info()
// 	if err != nil {
// 		log.Fatalf("Error getting response: %s", err)
// 	}
// 	// Check response status
// 	if res.IsError() {
// 		log.Fatalf("Error: %s", res.String())
// 	}
// 	// Deserialize the response into a map.
// 	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
// 		log.Fatalf("Error parsing the response body: %s", err)
// 	}
// 	// Print client and server version numbers.
// 	log.Printf("Client: %s", elasticsearch.Version)
// 	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
// 	log.Println(strings.Repeat("~", 37))

// 	// 2. Index documents concurrently
// 	//
// 	//for i, title := range []string{"Test One", "Test Two"} {
// 	//	wg.Add(1)
// 	//
// 	//	go func(i int, title string) {
// 	//		defer wg.Done()
// 	//
// 	//		// Build the request body.
// 	//		var b strings.Builder
// 	//		b.WriteString(`{"title" : "`)
// 	//		b.WriteString(title)
// 	//		b.WriteString(`"}`)
// 	//
// 	//		// Set up the request object.
// 	//		req := esapi.IndexRequest{
// 	//			Index:      "test",
// 	//			DocumentID: strconv.Itoa(i + 1),
// 	//			Body:       strings.NewReader(b.String()),
// 	//			Refresh:    "true",
// 	//		}
// 	//
// 	//		// Perform the request with the client.
// 	//		res, err := req.Do(context.Background(), es)
// 	//		if err != nil {
// 	//			log.Fatalf("Error getting response: %s", err)
// 	//		}
// 	//		defer res.Body.Close()
// 	//
// 	//		if res.IsError() {
// 	//			log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
// 	//		} else {
// 	//			// Deserialize the response into a map.
// 	//			var r map[string]interface{}
// 	//			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
// 	//				log.Printf("Error parsing the response body: %s", err)
// 	//			} else {
// 	//				// Print the response status and indexed document version.
// 	//				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
// 	//			}
// 	//		}
// 	//	}(i, title)
// 	//}
// 	//wg.Wait()
// 	//
// 	//log.Println(strings.Repeat("-", 37))

// 	// 3. Search for the indexed documents
// 	//
// 	// Build the request body.
// 	var buf bytes.Buffer
// 	query := map[string]interface{}{
// 		"query": map[string]interface{}{
// 			"match": map[string]interface{}{
// 				"title": "数据",
// 			},
// 		},
// 	}

// 	if err := json.NewEncoder(&buf).Encode(query); err != nil {
// 		log.Fatalf("Error encoding query: %s", err)
// 	}

// 	// Perform the search request.
// 	res, err = es.Search(
// 		es.Search.WithContext(context.Background()),
// 		es.Search.WithIndex("common_index_test_005"),
// 		//es.Search.WithSearchType("common_index_test_005"),
// 		es.Search.WithBody(&buf),
// 		es.Search.WithTrackTotalHits(true),
// 		es.Search.WithPretty(),
// 	)
// 	if err != nil {
// 		log.Fatalf("Error getting response: %s", err)
// 	}
// 	defer res.Body.Close()

// 	if res.IsError() {
// 		var e map[string]interface{}
// 		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
// 			log.Fatalf("Error parsing the response body: %s", err)
// 		} else {
// 			// Print the response status and error information.
// 			log.Fatalf("[%s] %s: %s",
// 				res.Status(),
// 				e["error"].(map[string]interface{})["type"],
// 				e["error"].(map[string]interface{})["reason"],
// 			)
// 		}
// 	}

// 	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
// 		log.Fatalf("Error parsing the response body: %s", err)
// 	}
// 	// Print the response status, number of results, and request duration.
// 	//log.Printf(
// 	//	"[%s] %d hits; took: %dms",
// 	//	res.Status(),
// 	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
// 	//	int(r["took"].(float64)),
// 	//)
// 	// Print the ID and document source for each hit.
// 	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
// 		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
// 	}

// 	log.Println(strings.Repeat("=", 37))
// }
