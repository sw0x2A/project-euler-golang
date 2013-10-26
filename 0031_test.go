package main

import "testing"

func TestP31(t *testing.T) {
	in, out := 200, 73682
	if x := p31(in); x != out {
		t.Errorf("p31(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP31(b *testing.B) {
	in := 200
	for i := 0; i < b.N; i++ {
		p31(in)
	}
}
