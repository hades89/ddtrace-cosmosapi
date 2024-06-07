package ddtrace

import (
	"context"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

const dbComponentName = "github.com/vippsas/go-cosmosdb/cosmosapi"

func StartSpanWithOperationName(ctx context.Context, operationName, resourceName string) (tracer.Span, context.Context) {
	opts := []ddtrace.StartSpanOption{
		tracer.ResourceName(resourceName),
		tracer.SpanType("azure-cosmosdb"),
		tracer.Tag(ext.DBType, "cosmosdb"),
		tracer.ServiceName("cosmosdb"),
		tracer.Tag(ext.Component, dbComponentName),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag("db.system", "cosmosdb"),
	}
	span, ctx := tracer.StartSpanFromContext(ctx, operationName, opts...)
	return span, ctx
}
