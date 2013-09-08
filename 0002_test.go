package main

import "testing"

func TestP2(t *testing.T) {
	in, out := 100, 44
	if x := p2(in); x != out {
		t.Errorf("p1(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP2(b *testing.B) {
	in := 1000
	for i := 0; i < b.N; i++ {
		p2(in)
	}
}
