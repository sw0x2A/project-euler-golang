package main

import "testing"

func TestP6(t *testing.T) {
	pairs := [...][2]uint{{10, 2640}, {100, 25164150}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p6(in); x != out {
			t.Errorf("p6(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP6(b *testing.B) {
	in := uint(100)
	for i := 0; i < b.N; i++ {
		p6(in)
	}
}
