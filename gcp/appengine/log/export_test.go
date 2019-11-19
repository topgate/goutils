package log

import (
	"io"
)

const (
	SeverityInfo     = severityInfo
	SeverityError    = severityError
	SeverityWarn     = severityWarn
	SeverityCritical = severityCritical
	SeverityDebug    = severityDebug
)

type Severify = severity

type CallerInfo = callerInfo

func SetCallerInfoGetter(fn func(skipFrame int) (CallerInfo, bool)) (finish func()) {
	old := getCallerInfo
	finish = func() { getCallerInfo = old }
	getCallerInfo = fn
	return
}

func SetWriterGetter(fn func() io.Writer) (finish func()) {
	old := getWriter
	finish = func() { getWriter = old }
	getWriter = fn
	return
}
