package main

import "testing"

func TestP1(t *testing.T) {
	in, out := 10, 23
	if x := p1(in); x != out {
		t.Errorf("p1(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP1(b *testing.B) {
	in := 1000
	for i := 0; i < b.N; i++ {
		p1(in)
	}
}
