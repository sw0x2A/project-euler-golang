package main

import "testing"

func TestP18(t *testing.T) {
	out := 1074
	if x := p18(); x != out {
		t.Errorf("p18() = %d, should be %d", x, out)
	}
}

func BenchmarkP18(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p18()
	}
}
