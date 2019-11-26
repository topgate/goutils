package validate

import "github.com/go-playground/validator/v10"

// Struct 与えられた構造体のバリデーション処理を行なう
func Struct(v interface{}) error {
	if err := validate.Struct(v); err != nil {
		ves := err.(validator.ValidationErrors)
		fes := make([]FieldError, 0, len(ves))
		for _, fe := range ves {
			fes = append(fes, fieldError{FieldError: fe})
		}
		return &ValidationError{FiledErrors: fes}
	}
	return nil
}
