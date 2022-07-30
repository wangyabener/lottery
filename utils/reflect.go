package utils

import (
	"reflect"
)

func InvokeObjectMethod(object interface{}, method string, args ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(object)

	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	ret := f.MethodByName(method).Call(inputs)
	return ret, nil
}
