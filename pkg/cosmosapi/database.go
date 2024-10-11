package cosmosapi

import (
	"context"
	"github.com/hades89/ddtrace-cosmosapi/pkg/ddtrace"
	"github.com/vippsas/go-cosmosdb/cosmosapi"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
)

func (t *TracedCosmosClient) CreateCollection(ctx context.Context, dbName string, colOps cosmosapi.CreateCollectionOptions) (cosmosapi.CreateCollectionResponse, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.create_collection", "CreateCollection", "", "")
	defer span.Finish()

	createCollectionResponse, err := t.client.CreateCollection(ctx, dbName, colOps)
	if err != nil {
		span.SetTag(ext.Error, err)
	}
	return createCollectionResponse, err
}

func (t *TracedCosmosClient) GetCollection(ctx context.Context, dbName, colName string) (*cosmosapi.Collection, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.get_collection", "GetCollection", colName, "")
	defer span.Finish()

	collection, err := t.client.GetCollection(ctx, dbName, colName)
	if err != nil {
		span.SetTag(ext.Error, err)
	}
	return collection, err
}

func (t *TracedCosmosClient) DeleteCollection(ctx context.Context, dbName, colName string) error {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.delete_collection", "DeleteCollection", colName, "")
	defer span.Finish()

	err := t.client.DeleteCollection(ctx, dbName, colName)
	if err != nil {
		span.SetTag(ext.Error, err)
	}
	return err
}

func (t *TracedCosmosClient) ReplaceCollection(ctx context.Context, dbName string, colOps cosmosapi.CollectionReplaceOptions) (*cosmosapi.Collection, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.replace_collection", "ReplaceCollection", "", "")
	defer span.Finish()

	collection, err := t.client.ReplaceCollection(ctx, dbName, colOps)
	if err != nil {
		span.SetTag(ext.Error, err)
	}
	return collection, err
}
