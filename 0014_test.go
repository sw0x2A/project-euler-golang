package main

import "testing"

func TestP14(t *testing.T) {
	in, out := 1000000, 837799
	if x := p14(in); x != out {
		t.Errorf("p14(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP14(b *testing.B) {
	in := 1000000
	for i := 0; i < b.N; i++ {
		p14(in)
	}
}
