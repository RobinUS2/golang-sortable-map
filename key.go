package sortablemap

import (
	"fmt"
	"reflect"
)

var supportedKeyTypes = map[reflect.Kind]bool{
	reflect.Int64:   true,
	reflect.Uint64:  true,
	reflect.Float64: true,
	reflect.String:  true,
}

func (k resultKey) Compare(keyType reflect.Kind, other resultKey) bool {
	if keyType == reflect.Int64 {
		v1 := k.Value.(int64)
		v2 := other.Value.(int64)
		return v1 < v2
	} else if keyType == reflect.Uint64 {
		v1 := k.Value.(uint64)
		v2 := other.Value.(uint64)
		return v1 < v2
	} else if keyType == reflect.Float64 {
		v1 := k.Value.(float64)
		v2 := other.Value.(float64)
		return v1 < v2
	} else if keyType == reflect.String {
		v1 := k.Value.(string)
		v2 := other.Value.(string)
		return v1 < v2
	} else {
		panic(fmt.Sprintf("comparator type %s not supported", keyType))
	}
}

type ResultKey struct {
	resultKey
	t reflect.Kind
}

func (k ResultKey) Int64() int64 {
	return k.Value.(int64)
}

func (k ResultKey) Uint64() uint64 {
	return k.Value.(uint64)
}

func (k ResultKey) String() string {
	if str, ok := k.Value.(string); ok {
		return str
	}
	return fmt.Sprintf("%+v", k.Value)
}

// internal key
type resultKey struct {
	Value interface{}
}
