package reflect

import (
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
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
