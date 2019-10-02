package log

type entity struct {
	Message        string               `json:"message"`
	Severity       severity             `json:"severity"`
	Trace          string               `json:"logging.googleapis.com/trace"`
	SourceLocation entitySourceLocation `json:"logging.googleapis.com/sourceLocation"`
}

type entitySourceLocation struct {
	File     string `json:"file"`
	Line     int    `json:"line"`
	Function string `json:"function"`
}
