package validate

import "github.com/go-playground/validator/v10"

// Struct 与えられた構造体のバリデーション処理を行なう
func Struct(v interface{}) error {
	if err := validate.Struct(v); err != nil {
		return validationErrors{ves: err.(validator.ValidationErrors)}
	}
	return nil
}
