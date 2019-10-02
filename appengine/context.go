package appengine

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"
)

type contextKey string

const (
	traceIDKey   contextKey = "trace_id"
	projectIDKey contextKey = "project_id"
)

// NewContext AppEngineのコンテキストを生成する
func NewContext(r *http.Request) context.Context {
	ctx := r.Context()

	if val := r.Header.Get("X-Cloud-Trace-Context"); val != "" {
		traceID := strings.SplitN(val, "/", 2)[0]
		ctx = context.WithValue(r.Context(), traceIDKey, traceID)
	} else {
		log.Println("cloud not get trace id from request header")
	}

	if projectID := os.Getenv("GOOGLE_CLOUD_PROJECT"); projectID != "" {
		ctx = context.WithValue(ctx, projectIDKey, projectID)
	} else {
		log.Println("cloud not get project id from environment variables")
	}
	return ctx
}

// TraceID ContextからトレースIDを取得する
func TraceID(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(traceIDKey).(string)
	return val, ok
}

// ProjectID ContextからプロジェクトIDを取得する
func ProjectID(ctx context.Context) (string, bool) {
	val, ok := ctx.Value(projectIDKey).(string)
	return val, ok
}
