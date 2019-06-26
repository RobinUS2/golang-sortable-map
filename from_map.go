package sortablemap

import (
	"errors"
	"fmt"
	"reflect"
)

// convenient function that directly gives iterator, may panic if types are not supported
func IteratorFromMap(v interface{}) (iterator *QueryResultIterator) {
	data, err := DataFromMap(v)
	if err != nil {
		panic(err)
	}
	return data.Iterator()
}

// use this function for proper error handling, can also be used to make multiple iterators of the same data error that's prepared (better performance)
func DataFromMap(v interface{}) (*Data, error) {
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
		return nil, fmt.Errorf("value type %s not supported", data.ValueType.String())
	}

	// prepare data set
	initResults := func(n int) {
		results = make(map[resultKey]resultValue, n)
	}
	addResult := func(k interface{}, v interface{}) {
		results[resultKey{
			Value: k,
		}] = resultValue{
			Value: v,
		}
	}

	// casts
	if data.KeyType == reflect.Int64 {
		if data.ValueType == reflect.Float64 {
			m := v.(map[int64]float64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Int64 {
			m := v.(map[int64]int64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Uint64 {
			m := v.(map[int64]uint64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Bool {
			m := v.(map[int64]bool)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Interface {
			m := v.(map[int64]interface{})
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		}
	} else if data.KeyType == reflect.String {
		if data.ValueType == reflect.Interface {
			m := v.(map[string]interface{})
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Float64 {
			m := v.(map[string]float64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Int64 {
			m := v.(map[string]int64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Uint64 {
			m := v.(map[string]uint64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Bool {
			m := v.(map[string]bool)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		}
	} else if data.KeyType == reflect.Uint64 {
		if data.ValueType == reflect.Float64 {
			m := v.(map[uint64]float64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Int64 {
			m := v.(map[uint64]int64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Uint64 {
			m := v.(map[uint64]uint64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Bool {
			m := v.(map[uint64]bool)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Interface {
			m := v.(map[uint64]interface{})
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		}
	} else if data.KeyType == reflect.Float64 {
		if data.ValueType == reflect.Float64 {
			m := v.(map[float64]float64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Bool {
			m := v.(map[float64]bool)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Int64 {
			m := v.(map[float64]int64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Uint64 {
			m := v.(map[float64]uint64)
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		} else if data.ValueType == reflect.Interface {
			m := v.(map[float64]interface{})
			initResults(len(m))
			for k, v := range m {
				addResult(k, v)
			}
		}
	}

	// validate
	if results == nil {
		panic(fmt.Sprintf("type not supported key %s values %s", data.KeyType, data.ValueType))
	}

	data.results = results
	return data, nil
}
