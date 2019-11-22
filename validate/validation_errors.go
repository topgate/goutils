package validate

import (
	"bytes"

	"github.com/go-playground/validator/v10"
)

// ValidationErrors go-playground/validatorから返される詳細なエラー
type ValidationErrors = validator.ValidationErrors

type validationErrors struct {
	ves ValidationErrors
}

func (v validationErrors) Error() string {
	transMessages := v.ves.Translate(Translator)
	builder := bytes.NewBufferString("")
	for _, m := range transMessages {
		builder.WriteString(m)
		builder.WriteString("\n")
	}
	builder.Truncate(builder.Len() - 1)
	return builder.String()
}

// GetValidationErrors go-playground/validatorから返される詳細なエラーを取得する
func GetValidationErrors(err error) (ValidationErrors, bool) {
	errs, ok := err.(validationErrors)
	return errs.ves, ok
}
