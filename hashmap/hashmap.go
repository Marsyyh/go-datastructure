package main

import (
	"fmt"
	"unsafe"
)

const (
	bucketCntBits = 3
	butcketCnt    = 1 << bucketCntBits
	loadFactor    = 6.5
)

type hmap struct {
	count     int            // map size
	B         uint8          // log_2 of # of buckets
	noverflow uint16         // approximate number of overfow buckets
	buckets   unsafe.Pointer // first buckets pointer for buckets array
}

type bmap struct {
	tophash [butcketCnt]uint8
}

// func NewTest(val int) *test {
// 	return &test{val}
// }

type Test struct {
	val int64
}

func (t *Test) size() int {
	size := int(unsafe.Sizeof(*t))
	return size
}

func main() {
	a := "test"
	t := &a
	d := &t
	fmt.Println(d, t, a)
	// t := reflect.TypeOf(a)
	// fmt.Println(t.Size())
}
