package sliceofpointers

import (
	"testing"
)

var resultValues []MyStruct
var resultPointers []*MyStruct

func BenchmarkSliceOfPointers(b *testing.B) {
	b.ReportAllocs()

	lenOfSlice := 20

	for i := 0; i < b.N; i++ {
		slice := make([]*MyStruct, 0, lenOfSlice)
		for j := 0; j < lenOfSlice; j++ {
			slice = append(slice, &MyStruct{A: j, B: j + 1})
		}

		resultPointers = slice
	}
}

func BenchmarkSliceOfValues(b *testing.B) {
	b.ReportAllocs()

	lenOfSlice := 20

	for i := 0; i < b.N; i++ {
		slice := make([]MyStruct, 0, lenOfSlice)
		for j := 0; j < lenOfSlice; j++ {
			slice = append(slice, MyStruct{A: j, B: j + 1})
		}

		resultValues = slice
	}
}
