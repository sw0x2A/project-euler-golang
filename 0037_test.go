package main

import "testing"

func TestP37(t *testing.T) {
	out := 748317
	if x := p37(); x != out {
		t.Errorf("p37() = %d, should be %d", x, out)
	}
}

func BenchmarkP37(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p37()
	}
}
