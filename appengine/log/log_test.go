package log_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"net/http"
	"os"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/topgate/goutils/appengine"
	. "github.com/topgate/goutils/appengine/log"
)

func TestLog(t *testing.T) {
	type (
		in struct {
			format string
			v      []interface{}
		}
		contextValue struct {
			projectID string
			traceID   string
		}
	)
	cases := []struct {
		name         string
		logFunc      func(context.Context, string, ...interface{})
		contextValue contextValue
		in           in
		expected     map[string]interface{}
	}{
		{
			name: "Debugf",
			in: in{
				format: "message: %v",
				v:      []interface{}{"value"},
			},
			logFunc: Debugf,
			contextValue: contextValue{
				projectID: "projectID",
				traceID:   "b86f2ca7a32d8569f517b6ccabc61e30/15904893325504441799;o=1",
			},
			expected: map[string]interface{}{
				"message":                      fmt.Sprintf("message: %v", "value"),
				"severity":                     "DEBUG",
				"logging.googleapis.com/trace": "projects/projectID/traces/b86f2ca7a32d8569f517b6ccabc61e30",
			},
		},
		{
			name: "Infof",
			in: in{
				format: "message: %v",
				v:      []interface{}{"value"},
			},
			logFunc: Infof,
			contextValue: contextValue{
				projectID: "projectID",
				traceID:   "b86f2ca7a32d8569f517b6ccabc61e30/15904893325504441799;o=1",
			},
			expected: map[string]interface{}{
				"message":                      fmt.Sprintf("message: %v", "value"),
				"severity":                     "INFO",
				"logging.googleapis.com/trace": "projects/projectID/traces/b86f2ca7a32d8569f517b6ccabc61e30",
			},
		},
		{
			name: "Warningf",
			in: in{
				format: "message: %v",
				v:      []interface{}{"value"},
			},
			logFunc: Waringf,
			contextValue: contextValue{
				projectID: "projectID",
				traceID:   "b86f2ca7a32d8569f517b6ccabc61e30/15904893325504441799;o=1",
			},
			expected: map[string]interface{}{
				"message":                      fmt.Sprintf("message: %v", "value"),
				"severity":                     "WARNING",
				"logging.googleapis.com/trace": "projects/projectID/traces/b86f2ca7a32d8569f517b6ccabc61e30",
			},
		},
		{
			name: "Errorf",
			in: in{
				format: "message: %v",
				v:      []interface{}{"value"},
			},
			logFunc: Errorf,
			contextValue: contextValue{
				projectID: "projectID",
				traceID:   "b86f2ca7a32d8569f517b6ccabc61e30/15904893325504441799;o=1",
			},
			expected: map[string]interface{}{
				"message":                      fmt.Sprintf("message: %v", "value"),
				"severity":                     "ERROR",
				"logging.googleapis.com/trace": "projects/projectID/traces/b86f2ca7a32d8569f517b6ccabc61e30",
			},
		},
		{
			name: "Criticalf",
			in: in{
				format: "message: %v",
				v:      []interface{}{"value"},
			},
			logFunc: Criticalf,
			contextValue: contextValue{
				projectID: "projectID",
				traceID:   "b86f2ca7a32d8569f517b6ccabc61e30/15904893325504441799;o=1",
			},
			expected: map[string]interface{}{
				"message":                      fmt.Sprintf("message: %v", "value"),
				"severity":                     "CRITICAL",
				"logging.googleapis.com/trace": "projects/projectID/traces/b86f2ca7a32d8569f517b6ccabc61e30",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertions := assert.New(t)

			buf := new(bytes.Buffer)
			finish := RewriteGetWriter(func() io.Writer { return buf })
			defer finish()

			ctx, ctxFinish := newContext(c.contextValue.projectID, c.contextValue.traceID)
			defer ctxFinish()
			c.logFunc(ctx, "message: %v", "value")

			var got map[string]interface{}
			if !assertions.NoError(json.Unmarshal(buf.Bytes(), &got)) {
				return
			}
			c.expected["logging.googleapis.com/sourceLocation"] = got["logging.googleapis.com/sourceLocation"]
			assertions.Equal(c.expected, got)
		})
	}
}

func newContext(projectID, traceID string) (ctx context.Context, finish func()) {
	os.Setenv("GOOGLE_CLOUD_PROJECT", projectID)
	r, _ := http.NewRequest("", "http://sample", nil)
	r.Header.Set("X-Cloud-Trace-Context", traceID)
	return appengine.NewContext(r), func() {
		os.Unsetenv("GOOGLE_CLOUD_PROJECT")
	}
}
