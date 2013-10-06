package main

import "testing"

func TestP24(t *testing.T) {
	in1, in2, out := 10, 1000000, 2783915460
	if x := p24(in1, in2); x != out {
		t.Errorf("p24(%d, %d) = %d, should be %d", in1, in2, x, out)
	}
}

func BenchmarkP24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p24(10, 1000000)
	}
}
