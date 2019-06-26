package sortablemap

import "reflect"

var supportedValueTypes = map[reflect.Kind]bool{
	reflect.Int64:     true,
	reflect.Float64:   true,
	reflect.Bool:      true,
	reflect.Ptr:       true, // pointer (e.g. *package.Type )
	reflect.Interface: true, // generic value without helper types
}

type ResultValue struct {
	resultValue
	t reflect.Kind
}

func (k ResultValue) Float64() float64 {
	return k.Value.(float64)
}

func (k ResultValue) Int64() int64 {
	return k.Value.(int64)
}

func (k ResultValue) Bool() bool {
	return k.Value.(bool)
}

func (k ResultValue) Interface() interface{} {
	return k.Value
}

func (k ResultValue) Pointer() *interface{} {
	return k.Value.(*interface{})
}

// internal value
type resultValue struct {
	Value interface{}
}
