package log

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"

	"github.com/topgate/goutils/appengine"
)

type severity string

const (
	severityDebug    severity = "DEBUG"
	severityInfo     severity = "INFO"
	severityError    severity = "ERROR"
	severityWarn     severity = "WARNING"
	severityCritical severity = "CRITICAL"
)

// Debugf デバックログを出力する
func Debugf(ctx context.Context, format string, v ...interface{}) {
	printf(ctx, severityDebug, format, v...)
}

// Infof 情報ログを出力する
func Infof(ctx context.Context, format string, v ...interface{}) {
	printf(ctx, severityInfo, format, v...)
}

// Waringf 警告ログを出力する
func Waringf(ctx context.Context, format string, v ...interface{}) {
	printf(ctx, severityWarn, format, v...)
}

// Errorf エラーログを出力する
func Errorf(ctx context.Context, format string, v ...interface{}) {
	printf(ctx, severityError, format, v...)
}

// Criticalf 重大ログを出力する
func Criticalf(ctx context.Context, format string, v ...interface{}) {
	printf(ctx, severityCritical, format, v...)
}

func printf(ctx context.Context, s severity, format string, v ...interface{}) {
	var (
		traceID, _        = appengine.TraceID(ctx)
		projectID, _      = appengine.ProjectID(ctx)
		pc, file, line, _ = runtime.Caller(2)
		f                 = runtime.FuncForPC(pc)
	)
	entry := entity{
		Message:  fmt.Sprintf(format, v...),
		Severity: s,
		Trace:    fmt.Sprintf("projects/%s/traces/%s", projectID, traceID),
		SourceLocation: entitySourceLocation{
			File:     file,
			Line:     line,
			Function: f.Name(),
		},
	}
	payload, _ := json.Marshal(entry)
	log.New(getWriter(), "", 0).Println(string(payload))
}

var getWriter = func() io.Writer {
	return os.Stdout
}
