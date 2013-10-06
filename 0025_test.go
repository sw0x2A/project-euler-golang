package main

import "testing"

func TestP25(t *testing.T) {
	in, out := 1000, 4782
	if x := p25(in); x != out {
		t.Errorf("p25(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP25(b *testing.B) {
	in := 1000
	for i := 0; i < b.N; i++ {
		p25(in)
	}
}
