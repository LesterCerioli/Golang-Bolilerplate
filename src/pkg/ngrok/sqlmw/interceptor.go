package sqlmw

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/ngrok/sqlmw"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/opentracing/opentracing-go/log"
	"io"
	"reflect"
	"strings"
)

type Interceptor struct {
	sqlmw.NullInterceptor
	Tracer opentracing.Tracer
}

// Rows interceptors
func (in Interceptor) RowsNext(ctx context.Context, rows driver.Rows, dest []driver.Value) (err error) {
	childSpan, err := in.childSpanFromContext(ctx, "RowsNext")
	if err != nil {
		return rows.Next(dest)
	}

	defer func() {
		if err != io.EOF {
			childSpan = in.spanError(childSpan, err)
		}
		childSpan.Finish()
	}()

	return rows.Next(dest)
}

// Stmt interceptors
func (in Interceptor) StmtClose(ctx context.Context, stmt driver.Stmt) (err error) {
	childSpan, _ := in.childSpanFromContext(ctx, "StmtClose")

	defer func() {
		if err != nil {
			childSpan = in.spanError(childSpan, err)
		}
		childSpan.Finish()
	}()

	return stmt.Close()
}

func (in Interceptor) StmtExecContext(ctx context.Context, stmt driver.StmtExecContext, query string, args []driver.NamedValue) (rows driver.Result, err error) {
	childSpan, _ := in.childSpanFromContext(ctx, "StmtExecContext")
	childSpan.SetTag("query", query)
	childSpan.SetTag("args", formatArgs(args))

	defer func() {
		if err != nil {
			childSpan = in.spanError(childSpan, err)
		}
		childSpan.Finish()
	}()

	return stmt.ExecContext(ctx, args)
}

func (in Interceptor) StmtQueryContext(ctx context.Context, stmt driver.StmtQueryContext, query string, args []driver.NamedValue) (rows driver.Rows, err error) {
	childSpan, _ := in.childSpanFromContext(ctx, "StmtQueryContext")
	childSpan.SetTag("query", query)
	childSpan.SetTag("args", formatArgs(args))

	defer func() {
		if err != nil {
			childSpan = in.spanError(childSpan, err)
		}
		childSpan.Finish()
	}()

	return stmt.QueryContext(ctx, args)
}

// Tx interceptors
func (in Interceptor) TxCommit(ctx context.Context, tx driver.Tx) (err error) {
	childSpan, _ := in.childSpanFromContext(ctx, "TxCommit")

	defer func() {
		if err != nil {
			childSpan = in.spanError(childSpan, err)
		}
		childSpan.Finish()
	}()

	return tx.Commit()
}

func (in Interceptor) TxRollback(ctx context.Context, tx driver.Tx) (err error) {
	childSpan, _ := in.childSpanFromContext(ctx, "TxRollback")

	defer func() {
		if err != nil {
			childSpan = in.spanError(childSpan, err)
		}
		childSpan.Finish()
	}()

	return tx.Rollback()
}

func (in Interceptor) spanError(span opentracing.Span, err error) opentracing.Span {
	ext.Error.Set(span, true)
	span.LogFields(
		log.String("event", "error"),
		log.String("message", err.Error()),
	)

	return span
}

func (in Interceptor) childSpanFromContext(ctx context.Context, operationName string) (opentracing.Span, error) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return nil, fmt.Errorf("Span not found in context")
	}

	childSpan := in.Tracer.StartSpan(
		operationName,
		opentracing.ChildOf(span.Context()),
	)
	childSpan.SetTag("component", "database/sql")

	return childSpan, nil
}

func formatArgs(args interface{}) string {
	argsVal := reflect.ValueOf(args)
	if argsVal.Kind() != reflect.Slice {
		return "<unknown>"
	}

	strArgs := make([]string, 0, argsVal.Len())
	for i := 0; i < argsVal.Len(); i++ {
		strArgs = append(strArgs, formatArg(argsVal.Index(i).Interface()))
	}

	return fmt.Sprintf("{%s}", strings.Join(strArgs, ", "))
}

func formatArg(arg interface{}) string {
	strArg := ""
	switch arg := arg.(type) {
	case []uint8:
		strArg = fmt.Sprintf("[%T len:%d]", arg, len(arg))
	case string:
		strArg = fmt.Sprintf("[%T %q]", arg, arg)
	case driver.NamedValue:
		if arg.Name != "" {
			strArg = fmt.Sprintf("[%T %s=%v]", arg.Value, arg.Name, formatArg(arg.Value))
		} else {
			strArg = formatArg(arg.Value)
		}
	default:
		strArg = fmt.Sprintf("[%T %v]", arg, arg)
	}

	return strArg
}
