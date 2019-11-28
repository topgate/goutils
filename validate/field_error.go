package validate

import (
	"github.com/go-playground/validator/v10"
)

// FieldError フィールドのエラー
type FieldError interface {
	validator.FieldError
	Error() string
}

type fieldError struct {
	validator.FieldError
}

func (fe fieldError) Error() string {
	return fe.Translate(translator)

}
