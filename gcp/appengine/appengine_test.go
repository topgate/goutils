package appengine_test

import (
	"net/http"
	"os"
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
	. "github.com/topgate/goutils/gcp/appengine"
)

func TestNewContext(t *testing.T) {
	type (
		prepare struct {
			projectID   *string
			serviceName *string
			version     *string
		}
		expected struct {
			projectID   string
			serviceName string
			version     string
		}
	)
	cases := []struct {
		name     string
		prepare  prepare
		expected expected
	}{
		{
			name: "環境変数設定済み_環境変数の正しく設定されること",
			prepare: prepare{
				projectID:   strToPtr("projectid"),
				serviceName: strToPtr("service"),
				version:     strToPtr("version"),
			},
			expected: expected{
				projectID:   "projectid",
				serviceName: "service",
				version:     "version",
			},
		},
		{
			name: "環境変数未設定_空文字が設定されること",
			prepare: prepare{
				projectID:   nil,
				serviceName: nil,
				version:     nil,
			},
			expected: expected{
				projectID:   "",
				serviceName: "",
				version:     "",
			},
		},
	}

	setEnvs := func(p prepare) {
		if val := p.projectID; val != nil {
			os.Setenv("GOOGLE_CLOUD_PROJECT", *val)
		}
		if val := p.serviceName; val != nil {
			os.Setenv("GAE_SERVICE", *val)
		}
		if val := p.version; val != nil {
			os.Setenv("GAE_VERSION", *val)
		}
	}
	unsetEnvs := func() {
		_ = os.Unsetenv("GOOGLE_CLOUD_PROJECT")
		_ = os.Unsetenv("GAE_SERVICE")
		_ = os.Unsetenv("GAE_VERSION")
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertions := assert.New(t)

			setEnvs(c.prepare)
			defer unsetEnvs()

			r, _ := http.NewRequest(http.MethodGet, "", nil)
			got := NewContext(r)

			assertions.Equal(c.expected.projectID, got.Value(ContextKeyProjectID))
			assertions.Equal(c.expected.serviceName, got.Value(ContextKeyServiceName))
			assertions.Equal(c.expected.version, got.Value(ContextKeyVersion))
		})
	}
}

func TestProjectID(t *testing.T) {
	cases := []struct {
		name     string
		in       context.Context
		expected string
	}{
		{
			name:     "プロジェクトIDの指定済み_設定されたプロジェクトID取得",
			in:       contextWithValue(ContextKeyProjectID, "projectid"),
			expected: "projectid",
		},
		{
			name:     "プロジェクトIDの未設定_空文字取得",
			in:       context.Background(),
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertions := assert.New(t)

			got := ProjectID(c.in)
			assertions.Equal(c.expected, got)
		})
	}
}

func TestServiceName(t *testing.T) {
	cases := []struct {
		name     string
		in       context.Context
		expected string
	}{
		{
			name:     "サービス名設定済み_設定されたサービス名取得",
			in:       contextWithValue(ContextKeyServiceName, "service"),
			expected: "service",
		},
		{
			name:     "サービス名の未設定_空文字取得",
			in:       context.Background(),
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertions := assert.New(t)

			got := ServiceName(c.in)
			assertions.Equal(c.expected, got)
		})
	}
}

func TestVersion(t *testing.T) {
	cases := []struct {
		name     string
		in       context.Context
		expected string
	}{
		{
			name:     "バージョン設定済み_設定されたバージョン取得",
			in:       contextWithValue(ContextKeyVersion, "version"),
			expected: "version",
		},
		{
			name:     "バージョンの未設定_空文字取得",
			in:       context.Background(),
			expected: "",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertions := assert.New(t)

			got := Version(c.in)
			assertions.Equal(c.expected, got)
		})
	}
}

func strToPtr(s string) *string {
	return &s
}

func contextWithValue(key, value interface{}) context.Context {
	return context.WithValue(context.Background(), key, value)
}
