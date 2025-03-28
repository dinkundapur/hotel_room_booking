package utils

import (
	"reflect"
)

// this functions accepts the struct object and call all the exported methods defined in the structs
func ExecuteMethods(obj interface{}) {
	// Calling all the struct methods
	structValue := reflect.ValueOf(obj)
	for i := 0; i < structValue.NumMethod(); i++ {
		method := structValue.Method(i)
		method.Call(nil)
	}
}
