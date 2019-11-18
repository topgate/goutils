package log_test

import (
	"bytes"
	"context"
	"io"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/topgate/goutils/gcp/appengine"
	. "github.com/topgate/goutils/gcp/appengine/log"
)

func TestPrintf(t *testing.T) {
	info := CallerInfo{
		File:   "log_test.go",
		FnName: "github.com/topgate/goutils/gcp/appengine/log.TestPrintf",
		Line:   10,
	}
	const (
		projectID = "printf_test"
		traceID   = "printf_test_trace_id"
	)

	type (
		in struct {
			format string
			v      []interface{}
		}
	)
	cases := []struct {
		name     string
		fn       func(ctx context.Context, format string, v ...interface{})
		in       in
		expected string
	}{
		{
			name: "Debugf",
			fn:   Debugf,
			in: in{
				format: "test %v",
				v:      []interface{}{"debug"},
			},
			expected: expectLogMessage(SeverityDebug, projectID, traceID, info.File, info.FnName, info.Line, fmt.Sprintf("test %v", "debug")),
		},
		{
			name: "Infof",
			fn:   Infof,
			in: in{
				format: "test2 %v",
				v:      []interface{}{"info"},
			},
			expected: expectLogMessage(SeverityInfo, projectID, traceID, info.File, info.FnName, info.Line, fmt.Sprintf("test2 %v", "info")),
		},
		{
			name: "Warningf",
			fn:   Warningf,
			in: in{
				format: "test3 %v",
				v:      []interface{}{"warning"},
			},
			expected: expectLogMessage(SeverityWarn, projectID, traceID, info.File, info.FnName, info.Line, fmt.Sprintf("test3 %v", "warning")),
		},
		{
			name: "Errorf",
			fn:   Errorf,
			in: in{
				format: "test4 %v",
				v:      []interface{}{"error"},
			},
			expected: expectLogMessage(SeverityError, projectID, traceID, info.File, info.FnName, info.Line, fmt.Sprintf("test4 %v", "error")),
		},
		{
			name: "Criticalf",
			fn:   Criticalf,
			in: in{
				format: "test5 %v",
				v:      []interface{}{"critical"},
			},
			expected: expectLogMessage(SeverityCritical, projectID, traceID, info.File, info.FnName, info.Line, fmt.Sprintf("test5 %v", "critical")),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var w bytes.Buffer

			wFinish := SetWriterGetter(func() io.Writer { return &w })
			cFinish := SetCallerInfoGetter(func(skipFrame int) (CallerInfo, bool) { return info, true })
			defer func() {
				wFinish()
				cFinish()
			}()

			assertions := assert.New(t)
			ctx := appengine.NewContext(setupEnvAndNewRequest(projectID, traceID))
			c.fn(ctx, c.in.format, c.in.v...)
			assertions.Equal(c.expected, w.String())
		})
	}
}

func TestPrint(t *testing.T) {
	info := CallerInfo{
		File:   "log_test.go",
		FnName: "github.com/topgate/goutils/gcp/appengine/log.TestPrint",
		Line:   10,
	}
	const (
		projectID = "print_test_test_project"
		traceID   = "print_test_trace_id"
	)

	cases := []struct {
		name     string
		fn       func(ctx context.Context, v interface{})
		in       interface{}
		expected string
	}{
		{
			name:     "Debug",
			fn:       Debug,
			in:       "test",
			expected: expectLogMessage(SeverityDebug, projectID, traceID, info.File, info.FnName, info.Line, "test"),
		},
		{
			name:     "Info",
			fn:       Info,
			in:       "test2",
			expected: expectLogMessage(SeverityInfo, projectID, traceID, info.File, info.FnName, info.Line, "test2"),
		},
		{
			name:     "Warning",
			fn:       Warning,
			in:       "test3",
			expected: expectLogMessage(SeverityWarn, projectID, traceID, info.File, info.FnName, info.Line, "test3"),
		},
		{
			name:     "Error",
			fn:       Error,
			in:       "test4",
			expected: expectLogMessage(SeverityError, projectID, traceID, info.File, info.FnName, info.Line, "test4"),
		},
		{
			name:     "Critical",
			fn:       Critical,
			in:       "test5",
			expected: expectLogMessage(SeverityCritical, projectID, traceID, info.File, info.FnName, info.Line, "test5"),
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var w bytes.Buffer

			wFinish := SetWriterGetter(func() io.Writer { return &w })
			cFinish := SetCallerInfoGetter(func(skipFrame int) (CallerInfo, bool) { return info, true })
			defer func() {
				wFinish()
				cFinish()
			}()

			assertions := assert.New(t)
			ctx := appengine.NewContext(setupEnvAndNewRequest(projectID, traceID))
			c.fn(ctx, c.in)
			assertions.Equal(c.expected, w.String())
		})
	}
}

func expectLogMessage(s Severify, projectID, traceID, fileName, funcName string, line int, msg string) string {
	return fmt.Sprintln(
		fmt.Sprintf(
			`{"logging.googleapis.com/sourceLocation":{"file":"%s","function":"%s","line":%d},"logging.googleapis.com/trace":"projects/%s/traces/%s","message":"%s","severity":"%s"}`,
			fileName, funcName, line, projectID, traceID, msg, s,
		),
	)
}
