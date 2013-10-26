package main

import "testing"

func TestP34(t *testing.T) {
	out := 40730
	if x := p34(); x != out {
		t.Errorf("p34() = %d, should be %d", x, out)
	}
}

func BenchmarkP34(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p34()
	}
}
