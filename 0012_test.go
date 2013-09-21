package main

import "testing"

func TestP12(t *testing.T) {
	out := 76576500
	if x := p12(); x != out {
		t.Errorf("p12() = %d, should be %d", x, out)
	}
}

func BenchmarkP12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p12()
	}
}
