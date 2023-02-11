package bench_test

import (
	"fmt"
	"main/bench"
	"testing"
)

var blackhole int

// func TestFileLen(t *testing.T) {
// 	result, err := bench.FileLen("testdata/american-english", 1)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if result != 985084 {
// 		t.Error("Expected 985084, got", result)
// 	}
// }

func BenchmarkFileLen1(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000, 1000000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := bench.FileLen("testdata/american-english", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
