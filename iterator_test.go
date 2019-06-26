package sortablemap_test

import (
	"fmt"
	"github.com/RobinUS2/golang-sortable-map"
	"testing"
)

func TestExample(t *testing.T) {
	iterator := sortablemap.IteratorFromMap(map[int64]float64{
		1: 2.0,
		4: 5.0,
		2: 3.0,
		0: 3.0,
	})
	for iterator.Next() {
		k, v := iterator.Value()
		fmt.Printf("%d %f\n", k.Int64(), v.Float64())
	}
}

func TestFromMap(t *testing.T) {
	m := map[int64]float64{
		1: 2.0,
		4: 5.0,
		2: 3.0,
		0: 3.0,
	}
	data, err := sortablemap.DataFromMap(m)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error()
	}
	if data.Len() != 4 {
		t.Error(data)
	}
	iter := data.Iterator()
	for iter.Next() {
		k, v := iter.Value()
		i := k.Int64()
		f := v.Float64()
		t.Logf("%+v %+v %d %.0f", k, v, i, f)
	}

	// clone and run again
	{
		cloned := iter.Clone()
		n := 0
		for cloned.Next() {
			n++
		}
		if n != 4 {
			t.Error(n)
		}
	}
}

func TestFromMapInterface(t *testing.T) {
	type testValueType struct {
		A float64
	}
	m := map[int64]interface{}{
		1: testValueType{2.0},
		4: testValueType{5.0},
		2: testValueType{3.0},
		0: testValueType{3.0},
	}
	data, err := sortablemap.DataFromMap(m)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error()
	}
	if data.Len() != 4 {
		t.Error(data)
	}
	iter := data.Iterator()
	for iter.Next() {
		k, v := iter.Value()
		i := k.Int64()
		f := v.Interface().(testValueType).A
		t.Logf("%+v %+v %d %.0f", k, v, i, f)
	}
}

func TestFromMapStringInterface(t *testing.T) {
	type testValueType struct {
		A float64
	}
	m := map[string]interface{}{
		"d": testValueType{2.0},
		"a": testValueType{5.0},
		"c": testValueType{3.0},
		"b": testValueType{3.0},
	}
	data, err := sortablemap.DataFromMap(m)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error()
	}
	if data.Len() != 4 {
		t.Error(data)
	}
	iter := data.Iterator()
	for iter.Next() {
		k, v := iter.Value()
		i := k.String()
		f := v.Interface().(testValueType).A
		t.Logf("%+v %+v %s %.0f", k, v, i, f)
	}
}

func TestFromMapFloatFloat(t *testing.T) {
	m := map[float64]float64{
		4.0: 2.0,
		1.0: 5.0,
		3.0: 3.0,
		2.0: 3.0,
	}
	data, err := sortablemap.DataFromMap(m)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error()
	}
	if data.Len() != 4 {
		t.Error(data)
	}
}

func TestFromMapFloatInt64(t *testing.T) {
	m := map[float64]int64{
		4.0: 2,
		1.0: 5,
		3.0: 3,
		2.0: 3,
	}
	data, err := sortablemap.DataFromMap(m)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error()
	}
	if data.Len() != 4 {
		t.Error(data)
	}
}

func TestFromMapUint64Ptr(t *testing.T) {
	type testPtrType struct {
		A int64
	}
	m := map[uint64]interface{}{
		4.0: &testPtrType{2},
		1.0: &testPtrType{5},
		3.0: &testPtrType{3},
		2.0: &testPtrType{3},
	}
	data, err := sortablemap.DataFromMap(m)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error()
	}
	if data.Len() != 4 {
		t.Error(data)
	}
}

func TestFromMapInt64Bool(t *testing.T) {
	m := map[int64]bool{
		4: true,
		2: true,
		3: true,
		1: true,
	}
	data, err := sortablemap.DataFromMap(m)
	if err != nil {
		t.Error(err)
	}
	if data == nil {
		t.Error()
	}
	if data.Len() != 4 {
		t.Error(data)
	}
	iter := data.Iterator()
	for iter.Next() {
		k, v := iter.Value()
		i := k.Int64()
		f := v.Bool()
		t.Logf("%+v %+v %d %v", k, v, i, f)
	}
}

func TestQueryResult_Iterator(t *testing.T) {
	res, err := sortablemap.DataFromMap(
		map[uint64]float64{
			1: 2.0,
			4: 5.0,
			2: 3.0,
			0: 3.0,
		},
	)
	if err != nil {
		t.Error(err)
	}
	n := res.Len()
	if n != 4 {
		t.Error(n)
	}

	// verify sorted
	iter := res.Iterator()
	{
		iterN := 0
		for iter.Next() {
			k, v := iter.Value()
			ts := k.Uint64()
			val := v.Float64()
			if iterN == 0 {
				if ts != 0 {
					t.Error(ts)
				}
				if val != 3.0 {
					t.Error()
				}
			} else if iterN == 1 {
				if ts != 1 {
					t.Error(ts)
				}
				if val != 2.0 {
					t.Error()
				}
			} else if iterN == 2 {
				if ts != 2 {
					t.Error(ts)
				}
				if val != 3.0 {
					t.Error()
				}
			} else if iterN == 3 {
				if ts != 4 {
					t.Error(ts)
				}
				if val != 5.0 {
					t.Error()
				}
			}
			iterN++
		}
		if iterN != n {
			t.Error(iterN, n)
		}
	}

	// call again, should be empty now
	if iter.Next() {
		t.Error()
	}

	// reset
	iter.Reset()

	// scan again
	{
		iterN := 0
		for iter.Next() {
			iterN++
		}
		if iterN != 4 {
			t.Error(iterN, 4)
		}
	}
}
