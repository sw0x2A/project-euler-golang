package main

import "testing"

func TestP32(t *testing.T) {
	out := 45228
	if x := p32(); x != out {
		t.Errorf("p32() = %d, should be %d", x, out)
	}
}

func BenchmarkP32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p32()
	}
}
