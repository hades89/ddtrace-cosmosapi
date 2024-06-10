package ddtrace

import (
	"context"
	"github.com/hades89/ddtrace-cosmosapi/pkg/common"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/ext"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func StartSpanWithOperationName(ctx context.Context, operationName, resourceName, collName, docId string) (tracer.Span, context.Context) {
	opts := []ddtrace.StartSpanOption{
		tracer.ServiceName(common.ServiceName),
		tracer.ResourceName(resourceName),
		tracer.SpanType(common.SpanType),
		tracer.Tag(ext.DBType, common.DbType),
		tracer.Tag(ext.Component, common.DbComponentName),
		tracer.Tag(ext.SpanKind, ext.SpanKindClient),
		tracer.Tag("db.system", common.DbSystem),
	}
	if collName != "" {
		opts = append(opts, tracer.Tag("cosmosdb.collection", collName))

	}
	if docId != "" {
		opts = append(opts, tracer.Tag("cosmosdb.document_id", docId))

	}
	span, ctx := tracer.StartSpanFromContext(ctx, operationName, opts...)
	return span, ctx
}