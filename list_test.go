package goj

import "testing"
import "fmt"

func testEach(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{1},
	}

	var times int
	fn := func() func(interface{}, int) {
		times = 0
		return func(v interface{}, k int) {
			fmt.Printf("Call function %d\n", k)
			times += 1
		}
	}

	for _, v := range cases {
		Each(fn(), []interface{}{v})
		if times != len(v) {
			t.Fatalf("Bad interation cycles: %d != %d", times, len(v))
		}
	}
}

func TestMap(t *testing.T) {
	cases := [][]int{
		{1, 2, 3},
		{1},
	}

	fn := func(n int) int {
		return n * 2
	}

	for _, v := range cases {
		r := Map(fn, v)
		if r[0] != 2 {
			t.Fatalf("Bad interation cycles: %d != %d", times, len(v))
		}
	}
}
