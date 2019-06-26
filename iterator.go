package sortablemap

import (
	"reflect"
	"sort"
)

type Data struct {
	results   map[resultKey]resultValue // in random order due to Go map implementation, if you need sorted data call Data.Iterator()
	KeyType   reflect.Kind
	ValueType reflect.Kind
}

func (res *Data) Len() int {
	return len(res.results)
}

func (res *Data) Iterator() *QueryResultIterator {
	iter := &QueryResultIterator{
		data: res,
		size: len(res.results),
	}
	iter.Reset()

	// sort
	sortedKeys := make([]resultKey, iter.size)
	idx := 0
	for k := range iter.data.results {
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

// clones iterator and resets all pointers
func (iter *QueryResultIterator) Clone() *QueryResultIterator {
	newIter := &QueryResultIterator{
		data:     &(*iter.data),
		current:  iter.current,
		size:     iter.size,
		dataKeys: iter.dataKeys,
	}
	newIter.Reset()
	return newIter
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
	value := iter.data.results[key]
	return ResultKey{key, iter.data.KeyType}, ResultValue{value, iter.data.ValueType}
}
