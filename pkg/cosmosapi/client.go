package cosmosapi

import "github.com/vippsas/go-cosmosdb/cosmosapi"

func NewTracedCosmosClient(client *cosmosapi.Client) *TracedCosmosClient {
	return &TracedCosmosClient{client: client}
}

type TracedCosmosClient struct {
	client *cosmosapi.Client
}
