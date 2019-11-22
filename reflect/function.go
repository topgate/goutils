// Package reflect リフレクション関係の機能をまとめたパッケージ
package reflect

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// GetFunctionName 関数名を取得する
func GetFunctionName(i interface{}) string {
	pationalNames := strings.Split(GetFunctionFullName(i), "/")
	return pationalNames[len(pationalNames)-1]

}

// GetFunctionFullName 関数のフルネームを取得する
func GetFunctionFullName(i interface{}) string {
	v := reflect.ValueOf(i)
	t := v.Type()
	if t.Kind() != reflect.Func {
		panic(fmt.Errorf("The %s kind is not function type", t.Kind()))
	}
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
