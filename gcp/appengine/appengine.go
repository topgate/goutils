package appengine

import (
	"context"
	"net/http"
	"os"
)

type contextKey string

const (
	contextKeyProjectID   contextKey = "project_id"
	contextKeyServiceName contextKey = "service_name"
	contextKeyVersion     contextKey = "version"
)

// NewContext App Engine用のContextオブジェクトを生成する
func NewContext(r *http.Request) context.Context {
	ctx := r.Context()

	ctx = context.WithValue(ctx, contextKeyProjectID, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	ctx = context.WithValue(ctx, contextKeyServiceName, os.Getenv("GAE_SERVICE"))
	ctx = context.WithValue(ctx, contextKeyVersion, os.Getenv("GAE_VERSION"))

	return ctx
}

// ProjectID GCPプロジェクトIDを取得する
// 値が取得できない場合は空文字列を返す
func ProjectID(ctx context.Context) string {
	return strOrBlank(ctx, contextKeyProjectID)
}

// ServiceName App Engineサービス名を取得する
// 値が取得できない場合は空文字列を返す
func ServiceName(ctx context.Context) string {
	return strOrBlank(ctx, contextKeyServiceName)
}

// Version App Engineサービス名を取得する
// 値が取得できない場合は空文字列を返す
func Version(ctx context.Context) string {
	return strOrBlank(ctx, contextKeyVersion)
}

func strOrBlank(ctx context.Context, key contextKey) string {
	val, ok := ctx.Value(key).(string)
	if !ok {
		return ""
	}
	return val
}
