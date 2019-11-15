package log

import (
	"net/http"
	"os"

	"github.com/topgate/goutils/gcp/appengine"
)

func Example() {
	r := setupEnvAndNewRequest("sample_project_id", "sample_trace_id/xxxxxxxxxx")
	ctx := appengine.NewContext(r)

	Debugf(ctx, "debug %v", "foo")
	Infof(ctx, "info %v", "foo")
	Warningf(ctx, "warning %v", "foo")
	Errorf(ctx, "error %v", "foo")
	Criticalf(ctx, "critical %v", "foo")

	// Output:
	// {"logging.googleapis.com/sourceLocation":{"file":"log_test.go","function":"github.com/topgate/goutils/gcp/appengine/log.Example","line":14},"logging.googleapis.com/trace":"projects/sample_project_id/traces/sample_trace_id","message":"debug foo","severity":"DEBUG"}
	// {"logging.googleapis.com/sourceLocation":{"file":"log_test.go","function":"github.com/topgate/goutils/gcp/appengine/log.Example","line":15},"logging.googleapis.com/trace":"projects/sample_project_id/traces/sample_trace_id","message":"info foo","severity":"INFO"}
	// {"logging.googleapis.com/sourceLocation":{"file":"log_test.go","function":"github.com/topgate/goutils/gcp/appengine/log.Example","line":16},"logging.googleapis.com/trace":"projects/sample_project_id/traces/sample_trace_id","message":"warning foo","severity":"WARNING"}
	// {"logging.googleapis.com/sourceLocation":{"file":"log_test.go","function":"github.com/topgate/goutils/gcp/appengine/log.Example","line":17},"logging.googleapis.com/trace":"projects/sample_project_id/traces/sample_trace_id","message":"error foo","severity":"ERROR"}
	// {"logging.googleapis.com/sourceLocation":{"file":"log_test.go","function":"github.com/topgate/goutils/gcp/appengine/log.Example","line":18},"logging.googleapis.com/trace":"projects/sample_project_id/traces/sample_trace_id","message":"critical foo","severity":"CRITICAL"}
}

func setupEnvAndNewRequest(projectID, traceID string) *http.Request {
	os.Setenv("GOOGLE_CLOUD_PROJECT", projectID)
	req, _ := http.NewRequest(http.MethodGet, "", nil)
	req.Header.Set("X-Cloud-Trace-Context", traceID)
	return req
}
