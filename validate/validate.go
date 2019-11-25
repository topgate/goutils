// Package validate 日本語環境向けバリデーションを行なうライブラリがあるパッケージ
package validate

import (
	"reflect"

	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
)

var (
	validate *validator.Validate
	// Translator validation結果メッセージの翻訳器
	Translator ut.Translator
)

func init() {
	ja := ja.New()
	uni := ut.New(ja, ja)
	validate = validator.New()
	var found bool
	Translator, found = uni.GetTranslator("ja")
	if !found {
		panic("translator not found")
	}
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		fieldName := fld.Tag.Get("fieldname")
		if fieldName == "-" {
			return ""
		}
		return fieldName
	})
	if err := ja_translations.RegisterDefaultTranslations(validate, Translator); err != nil {
		panic(err)
	}
}
