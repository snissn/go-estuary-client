package main

import (
	"context"
	"fmt"

	estuary_client "github.com/application-research/estuary-clients/go"
)

func main() {

	auth := context.WithValue(context.Background(), estuary_client.ContextAPIKey, estuary_client.APIKey{
		Key:    "ESTxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxARY",
		Prefix: "Bearer",
	})
	cfg := estuary_client.NewConfiguration()
	r := estuary_client.NewAPIClient(cfg)
	data, http_response, err := r.CollectionsApi.CollectionsGet(auth)
	if err != nil {
		panic(err)
	}
	_ = http_response
	fmt.Printf("%+v", http_response)
	fmt.Println("hi")
	fmt.Printf("%+v", data)
}
