package main

import "testing"

func TestP33(t *testing.T) {
	out := 100
	if x := p33(); x != out {
		t.Errorf("p33() = %d, should be %d", x, out)
	}
}

func BenchmarkP33(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p33()
	}
}
