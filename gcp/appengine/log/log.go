package log

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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
	debugf(ctx, fmt.Sprintf(format, v...))
}

// Debug デバッグログを出力する
func Debug(ctx context.Context, v interface{}) {
	debugf(ctx, fmt.Sprint(v))
}

func debugf(ctx context.Context, msg string) {
	logPrintf(ctx, severityDebug, 2, msg)
}

// Infof 情報ログを出力する
func Infof(ctx context.Context, format string, v ...interface{}) {
	infof(ctx, fmt.Sprintf(format, v...))
}

// Info 情報ログを出力する
func Info(ctx context.Context, v interface{}) {
	infof(ctx, fmt.Sprint(v))
}

func infof(ctx context.Context, msg string) {
	logPrintf(ctx, severityInfo, 2, msg)
}

// Errorf エラーログを出力する
func Errorf(ctx context.Context, format string, v ...interface{}) {
	errorf(ctx, fmt.Sprintf(format, v...))
}

// Error エラーログを出力する
func Error(ctx context.Context, v interface{}) {
	errorf(ctx, fmt.Sprint(v))
}

func errorf(ctx context.Context, msg string) {
	logPrintf(ctx, severityError, 2, msg)
}

// Warningf 警告ログを出力する
func Warningf(ctx context.Context, format string, v ...interface{}) {
	warningf(ctx, fmt.Sprintf(format, v...))
}

// Warning 警告ログを出力する
func Warning(ctx context.Context, v interface{}) {
	warningf(ctx, fmt.Sprint(v))
}

func warningf(ctx context.Context, msg string) {
	logPrintf(ctx, severityWarn, 2, msg)
}

// Criticalf 重大エラーログを出力する
func Criticalf(ctx context.Context, format string, v ...interface{}) {
	criticalf(ctx, fmt.Sprintf(format, v...))
}

// Critical 重大エラーログを出力する
func Critical(ctx context.Context, v interface{}) {
	criticalf(ctx, fmt.Sprint(v))
}

func criticalf(ctx context.Context, msg string) {
	logPrintf(ctx, severityCritical, 2, msg)
}

func logPrintf(ctx context.Context, s severity, skipFrame int, msg string) {
	// プレフィックスに余計な文字列がつかないようにLoggerオブジェクトを作成
	logger := log.New(getWriter(), "", 0)

	// 構造化ロギングに出力する情報
	var (
		projectID = appengine.ProjectID(ctx)
		traceID   = appengine.TraceID(ctx)
		info, _   = getCallerInfo(skipFrame + 1)
	)
	if projectID == "" || traceID == "" {
		logger.Println("application might not be run in App Engine")
	}

	// 設定可能な特殊フィールドについては次を参照
	//   https://cloud.google.com/logging/docs/agent/configuration?hl=ja#special-fields
	entry := map[string]interface{}{
		"message":                      msg,
		"severity":                     s,
		"logging.googleapis.com/trace": fmt.Sprintf("projects/%s/traces/%s", projectID, traceID),
		"logging.googleapis.com/sourceLocation": map[string]interface{}{
			"file":     info.File,
			"line":     info.Line,
			"function": info.FnName,
		},
	}
	payload, _ := json.Marshal(entry)
	logger.Println(string(payload))
}

var getWriter = func() io.Writer {
	return os.Stdout
}

type callerInfo struct {
	File   string
	Line   int
	FnName string
}

var getCallerInfo = func(skipFrame int) (callerInfo, bool) {
	pc, file, line, ok := runtime.Caller(skipFrame + 1)
	if !ok {
		return callerInfo{}, ok
	}
	f := runtime.FuncForPC(pc)
	return callerInfo{
		File:   path.Base(file),
		Line:   line,
		FnName: f.Name(),
	}, true
}
