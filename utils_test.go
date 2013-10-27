package main

import "testing"

type primeData struct {
	in  int
	out []int
}

func TestUtilsPrimesBelow(t *testing.T) {
	pairs := []primeData{
		{0, []int{}},
		{1, []int{}},
		{2, []int{2}},
		{3, []int{2, 3}},
		{4, []int{2, 3}},
		{5, []int{2, 3, 5}},
		{6, []int{2, 3, 5}},
		{7, []int{2, 3, 5, 7}},
		{8, []int{2, 3, 5, 7}},
		{10, []int{2, 3, 5, 7}},
	}
	for _, pair := range pairs {
		in, out := pair.in, pair.out
		if x := PrimesBelow(in); !IntArrayEquals(x, out) {
			t.Errorf("PrimesBelow(%d) = %d, should be %d", in, x, out)
		}
	}
}

func TestUtilsPrimesBelowLength(t *testing.T) {
	pairs := [...][2]int{
		{10, 4},
		{100, 25},
		{1000, 168},
		{10000, 1229},
		{100000, 9592},
		{1000000, 78498},
		//{10000000, 664579},
		//{100000000, 5761455},
		//{1000000000, 50847534},
	}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := len(PrimesBelow(in)); x != out {
			t.Errorf("len(PrimesBelow(%d)) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkUtilsPrimesBelow(b *testing.B) {
	in := 4000000
	for i := 0; i < b.N; i++ {
		PrimesBelow(in)
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
