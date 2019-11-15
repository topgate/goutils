package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"

	"path"

	"github.com/topgate/goutils/gcp/appengine"
)

type severity string

// 利用可能なログレベルについては次を参照
//   https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry?hl=ja#LogSeverity
const (
	severityInfo     severity = "INFO"
	severityError    severity = "ERROR"
	severityWarn     severity = "WARNING"
	severityCritical severity = "CRITICAL"
	severityDebug    severity = "DEBUG"
)

// Debugf デバッグログを出力する
func Debugf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityDebug, format, v...)
}

// Infof 情報ログを出力する
func Infof(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityInfo, format, v...)
}

// Errorf エラーログを出力する
func Errorf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityError, format, v...)
}

// Warningf 警告ログを出力する
func Warningf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityWarn, format, v...)
}

// Criticalf 重大エラーログを出力する
func Criticalf(ctx context.Context, format string, v ...interface{}) {
	logPrintf(ctx, severityCritical, format, v...)
}

func logPrintf(ctx context.Context, s severity, format string, v ...interface{}) {
	// プレフィックスに余計な文字列がつかないようにLoggerオブジェクトを作成
	logger := log.New(os.Stdout, "", 0)

	// 構造化ロギングに出力する情報
	var (
		projectID         = appengine.ProjectID(ctx)
		traceID           = appengine.TraceID(ctx)
		pc, file, line, _ = runtime.Caller(2)
		f                 = runtime.FuncForPC(pc)
	)
	if projectID == "" || traceID == "" {
		logger.Println("application might not be run in App Engine")
	}

	// 設定可能な特殊フィールドについては次を参照
	//   https://cloud.google.com/logging/docs/agent/configuration?hl=ja#special-fields
	entry := map[string]interface{}{
		"message":                      fmt.Sprintf(format, v...),
		"severity":                     s,
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", projectID, traceID),
		"logging.googleapis.com/sourceLocation": map[string]interface{}{
			"file":     path.Base(file),
			"line":     line,
			"function": f.Name(),
		},
	}
	payload, _ := json.Marshal(entry)
	logger.Println(string(payload))
}
