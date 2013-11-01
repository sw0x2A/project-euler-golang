package main

import "testing"

func TestP39(t *testing.T) {
	out := 840
	if x := p39(); x != out {
		t.Errorf("p39() = %d, should be %d", x, out)
	}
}

func BenchmarkP39(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p39()
	}
}
