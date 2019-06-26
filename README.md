# golang-sortable-map  [![Build Status](https://travis-ci.org/RobinUS2/golang-sortable-map.svg?branch=master)](https://travis-ci.org/RobinUS2/tsxdb)
Map which can be iterated over in a sorted stable fashion

## Usage
```
import "github.com/RobinUS2/golang-sortable-map"
```

```
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

// outputs:
// 0 3.000000
// 1 2.000000
// 2 3.000000
// 4 5.000000
```

## Type support
Supported types right now are:

| key \ value 	| float64 	| bool 	| interface{} 	| int64 	| uint64 	|
|-------------	|---------	|------	|-------------	|-------	|--------	|
| float64     	| Y       	| Y    	| Y           	| Y     	| Y      	|
| int64       	| Y       	| Y    	| Y           	| Y     	| Y      	|
| uint64      	| Y       	| Y    	| Y           	| Y     	| Y      	|
| string      	| Y       	| Y    	| Y           	| Y     	| Y      	|