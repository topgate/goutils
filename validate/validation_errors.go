package validate

import (
	"bytes"
)

// ValidationError バリデーション時のエラー
type ValidationError struct {
	// FieldError フィールドごとのエラー
	FiledErrors []FieldError
}

func (v ValidationError) Error() string {
	builder := bytes.NewBufferString("")
	for _, m := range v.FiledErrors {
		builder.WriteString(m.Error())
		builder.WriteString("\n")
	}
	builder.Truncate(builder.Len() - 1)
	return builder.String()
}
