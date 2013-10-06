package main

import "testing"

func TestP26(t *testing.T) {
	pairs := [...][2]int{{10, 7}, {1000, 983}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p26(in); x != out {
			t.Errorf("p26(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP26(b *testing.B) {
	in := 1000
	for i := 0; i < b.N; i++ {
		p26(in)
	}
}
