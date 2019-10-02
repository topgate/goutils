package appengine

import (
	"fmt"
	"net/http"
	"os"
)

func Example() {
	// 環境変数の設定、リクエストの生成はHTTPハンドラー内で実際に利用するときは不要
	os.Setenv("GOOGLE_CLOUD_PROJECT", "sample_project_id")
	r, _ := http.NewRequest(http.MethodGet, "http://sample.com", nil)
	r.Header.Set("X-Cloud-Trace-Context", "b86f2ca7a32d8569f517b6ccabc61e30/15904893325504441799;o=1")

	ctx := NewContext(r)

	projectID, _ := ProjectID(ctx)
	traceID, _ := TraceID(ctx)
	fmt.Println(projectID)
	fmt.Println(traceID)

	// Output:
	// sample_project_id
	// b86f2ca7a32d8569f517b6ccabc61e30
}
