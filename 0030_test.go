package main

import "testing"

func TestP30(t *testing.T) {
	in, out := 5, 443839
	if x := p30(in); x != out {
		t.Errorf("p30(%d) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP30(b *testing.B) {
	in := 5
	for i := 0; i < b.N; i++ {
		p30(in)
	}
}
