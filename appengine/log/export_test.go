package log

import (
	"io"
)

func RewriteGetWriter(f func() io.Writer) (finish func()) {
	before := getWriter
	getWriter = f
	return func() {
		getWriter = before
	}
}
