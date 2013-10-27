package main

import "testing"

func TestP27(t *testing.T) {
	out := -59231
	if x := p27(); x != out {
		t.Errorf("p27() = %d, should be %d", x, out)
	}
}

func BenchmarkP27(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p27()
	}
}
