package cosmosapi

import (
	"context"
	"github.com/hades89/ddtrace-cosmosapi/pkg/ddtrace"
	"github.com/vippsas/go-cosmosdb/cosmosapi"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
)

func (t *TracedCosmosClient) GetDocument(ctx context.Context, dbName, collName, docID string, opts cosmosapi.GetDocumentOptions, out interface{}) (cosmosapi.DocumentResponse, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.get_document", "GetDocument", collName, docID)
	defer span.Finish()

	resp, err := t.client.GetDocument(ctx, dbName, collName, docID, opts, out)
	if err != nil {
		span.SetTag(ext.Error, err)
	}

	return resp, err
}

func (t *TracedCosmosClient) QueryDocuments(ctx context.Context, dbName, collName string, query cosmosapi.Query, out interface{}, opts cosmosapi.QueryDocumentsOptions) (cosmosapi.QueryDocumentsResponse, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.query_documents", "QueryDocuments", collName, "")
	defer span.Finish()

	resp, err := t.client.QueryDocuments(ctx, dbName, collName, query, out, opts)
	if err != nil {
		span.SetTag(ext.Error, err)
	}

	return resp, err
}

func (t *TracedCosmosClient) CreateDocument(ctx context.Context, dbName, collName string, doc interface{}, opts cosmosapi.CreateDocumentOptions) (*cosmosapi.Resource, cosmosapi.DocumentResponse, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.create_document", "CreateDocument", collName, "")
	defer span.Finish()

	resource, resp, err := t.client.CreateDocument(ctx, dbName, collName, doc, opts)
	if err != nil {
		span.SetTag(ext.Error, err)
	}

	return resource, resp, err
}

func (t *TracedCosmosClient) ReplaceDocument(ctx context.Context, dbName, collName, docID string, doc interface{}, opts cosmosapi.ReplaceDocumentOptions) (*cosmosapi.Resource, cosmosapi.DocumentResponse, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.replace_document", "ReplaceDocument", collName, docID)
	defer span.Finish()

	resource, resp, err := t.client.ReplaceDocument(ctx, dbName, collName, docID, doc, opts)
	if err != nil {
		span.SetTag(ext.Error, err)
	}

	return resource, resp, err
}

func (t *TracedCosmosClient) DeleteDocument(ctx context.Context, dbName, collName, docID string, opts cosmosapi.DeleteDocumentOptions) (cosmosapi.DocumentResponse, error) {
	span, ctx := ddtrace.StartSpanWithOperationName(ctx, "cosmosdb.delete_document", "DeleteDocument", collName, docID)
	defer span.Finish()

	resp, err := t.client.DeleteDocument(ctx, dbName, collName, docID, opts)
	if err != nil {
		span.SetTag(ext.Error, err)
	}

	return resp, err
}
