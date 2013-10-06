package main

import "testing"

func TestP21(t *testing.T) {
	pairs := [...][2]int{{10000, 31626}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p21(in); x != out {
			t.Errorf("p21(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP21(b *testing.B) {
	in := 10000
	for i := 0; i < b.N; i++ {
		p21(in)
	}
}
