package main

import "testing"

func TestP17(t *testing.T) {
	out := 21124
	if x := p17(); x != out {
		t.Errorf("p17() = %d, should be %d", x, out)
	}
}

func BenchmarkP17(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p17()
	}
}
