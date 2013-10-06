package main

import "testing"

func TestP20(t *testing.T) {
	pairs := [...][2]int{{10, 27}, {100, 648}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p20(in); x != out {
			t.Errorf("p20(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP20(b *testing.B) {
	in := 100
	for i := 0; i < b.N; i++ {
		p20(in)
	}
}
