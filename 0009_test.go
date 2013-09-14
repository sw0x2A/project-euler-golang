package main

import "testing"

func TestP9(t *testing.T) {
	in, out := 1000, 31875000
	if x := p9(in); x != out {
		t.Errorf("p9(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP9(b *testing.B) {
	in := 1000
	for i := 0; i < b.N; i++ {
		p9(in)
	}
}
