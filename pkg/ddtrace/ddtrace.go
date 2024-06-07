package ddtrace

import (
	"context"
	"github.com/hades89/ddtrace-cosmosapi/pkg/common"
	"github.com/vippsas/go-cosmosdb/cosmosapi"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func NewTracedCosmosClient(client *cosmosapi.Client) *TracedCosmosClient {
	return &TracedCosmosClient{client: client}
}

type TracedCosmosClient struct {
	client *cosmosapi.Client
}

func StartSpanWithOperationName(ctx context.Context, operationName, resourceName string) (tracer.Span, context.Context) {
	opts := []ddtrace.StartSpanOption{
		tracer.ServiceName(common.ServiceName),
		tracer.ResourceName(resourceName),
		tracer.SpanType(common.SpanType),
		tracer.Tag(ext.DBType, common.DbType),
		tracer.Tag(ext.Component, common.DbComponentName),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag("db.system", common.DbSystem),
	}
	span, ctx := tracer.StartSpanFromContext(ctx, operationName, opts...)
	return span, ctx
}
