package sortablemap_test

import (
	"github.com/RobinUS2/golang-sortable-map"
	"testing"
)

func TestFromMap(t *testing.T) {
	m := map[int64]float64{
		1: 2.0,
		4: 5.0,
		2: 3.0,
		0: 3.0,
	}
	data, err := sortablemap.FromMap(m)
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
}

func TestQueryResult_Iterator(t *testing.T) {
	res, err := sortablemap.FromMap(
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
	n := len(res.Results)
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
