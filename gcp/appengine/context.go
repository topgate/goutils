package appengine

import (
	"context"
	"net/http"
	"os"
	"strings"
)

type (
	contextKeyProjectID   struct{}
	contextKeyServiceName struct{}
	contextKeyVersion     struct{}
	contextKeyTraceID     struct{}
)

// NewContext App Engine用のContextオブジェクトを生成する
func NewContext(r *http.Request) context.Context {
	ctx := r.Context()

	ctx = context.WithValue(ctx, contextKeyProjectID{}, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	ctx = context.WithValue(ctx, contextKeyServiceName{}, os.Getenv("GAE_SERVICE"))
	ctx = context.WithValue(ctx, contextKeyVersion{}, os.Getenv("GAE_VERSION"))
	traceID := strings.SplitN(r.Header.Get("X-Cloud-Trace-Context"), "/", 2)[0]
	ctx = context.WithValue(ctx, contextKeyTraceID{}, traceID)

	return ctx
}

// ProjectID GCPプロジェクトIDを取得する
// 値が取得できない場合は空文字列を返す
func ProjectID(ctx context.Context) string {
	return strOrBlank(ctx, contextKeyProjectID{})
}

// ServiceName App Engineサービス名を取得する
// 値が取得できない場合は空文字列を返す
func ServiceName(ctx context.Context) string {
	return strOrBlank(ctx, contextKeyServiceName{})
}

// Version App Engineサービス名を取得する
// 値が取得できない場合は空文字列を返す
func Version(ctx context.Context) string {
	return strOrBlank(ctx, contextKeyVersion{})
}

// TraceID トレースIDの情報を取得する
// 値が取得できない場合は空文字列を返す
func TraceID(ctx context.Context) string {
	return strOrBlank(ctx, contextKeyTraceID{})
}

func strOrBlank(ctx context.Context, key interface{}) string {
	val, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return val
}
