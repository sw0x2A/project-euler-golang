package main

import "testing"

func TestP23(t *testing.T) {
	out := 4179871
	if x := p23(); x != out {
		t.Errorf("p23() = %d, should be %d", x, out)
	}
}

func BenchmarkP23(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p23()
	}
}
