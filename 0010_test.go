package main

import "testing"

func TestP10(t *testing.T) {
	in, out := 2000000, 142913828922
	if x := p10(in); x != out {
		t.Errorf("p10(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP10(b *testing.B) {
	in := 2000000
	for i := 0; i < b.N; i++ {
		p10(in)
	}
}
