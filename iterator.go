package sortablemap

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

var supportedKeyTypes = map[reflect.Kind]bool{
	reflect.Int64:  true,
	reflect.Uint64: true,
}

var supportedValueTypes = map[reflect.Kind]bool{
	reflect.Float64: true,
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

type ResultValue struct {
	resultValue
	t reflect.Kind
}

func (k ResultValue) Float64() float64 {
	return k.Value.(float64)
}

// internal key
type resultKey struct {
	Value interface{}
}

// internal value
type resultValue struct {
	Value interface{}
}

func FromMap(v interface{}) (*Data, error) {
	typeInfo := reflect.TypeOf(v)
	if typeInfo.Kind() != reflect.Map {
		return nil, errors.New("must be map")
	}
	var results map[resultKey]resultValue
	data := &Data{
		KeyType:   typeInfo.Key().Kind(),
		ValueType: typeInfo.Elem().Kind(),
	}

	// check types
	if !supportedKeyTypes[data.KeyType] {
		return nil, fmt.Errorf("key type %s not supported", data.KeyType.String())
	}
	if !supportedValueTypes[data.ValueType] {
		return nil, fmt.Errorf("value type %s not supported", data.KeyType.String())
	}

	// prepare data set
	initResults := func(n int) {
		results = make(map[resultKey]resultValue, n)
	}

	// casts
	if data.KeyType == reflect.Int64 && data.ValueType == reflect.Float64 {
		m := v.(map[int64]float64)
		initResults(len(m))
		for k, v := range m {
			results[resultKey{
				Value: k,
			}] = resultValue{
				Value: v,
			}
		}
	} else if data.KeyType == reflect.Uint64 && data.ValueType == reflect.Float64 {
		m := v.(map[uint64]float64)
		initResults(len(m))
		for k, v := range m {
			results[resultKey{
				Value: k,
			}] = resultValue{
				Value: v,
			}
		}
	} else {
		panic(fmt.Sprintf("type not supported key %s values %s", data.KeyType, data.ValueType))
	}
	if results == nil {
		panic("results empty")
	}
	data.Results = results
	return data, nil
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
	} else {
		panic(fmt.Sprintf("comparator type %s not supported", keyType))
	}
}

// @todo refactor to private values since modification will mess up things
type Data struct {
	Results   map[resultKey]resultValue // in random order due to Go map implementation, if you need sorted data call Data.Iterator()
	KeyType   reflect.Kind
	ValueType reflect.Kind
}

func (res *Data) Len() int {
	return len(res.Results)
}

func (res *Data) Iterator() *QueryResultIterator {
	iter := &QueryResultIterator{
		data: res,
		size: len(res.Results),
	}
	iter.Reset()

	// sort
	sortedKeys := make([]resultKey, iter.size)
	idx := 0
	for k := range iter.data.Results {
		sortedKeys[idx] = k
		idx++
	}
	sort.Slice(sortedKeys, func(i, j int) bool { return sortedKeys[i].Compare(res.KeyType, sortedKeys[j]) })
	iter.dataKeys = sortedKeys

	return iter
}

type QueryResultIterator struct {
	data     *Data
	current  int
	size     int
	dataKeys []resultKey
}

func (iter *QueryResultIterator) Reset() {
	iter.current = -1 // start before first, since you do for it.Next() { it.Value() }
}

func (iter *QueryResultIterator) Next() bool {
	iter.current++
	return iter.current < iter.size
}

func (iter *QueryResultIterator) Value() (ResultKey, ResultValue) {
	key := iter.dataKeys[iter.current]
	value := iter.data.Results[key]
	return ResultKey{key, iter.data.KeyType}, ResultValue{value, iter.data.ValueType}
}
