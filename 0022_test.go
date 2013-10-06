package main

import "testing"

func TestP22(t *testing.T) {
	out := 871198282
	if x := p22(); x != out {
		t.Errorf("p22() = %d, should be %d", x, out)
	}
}

func BenchmarkP22(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p22()
	}
}
