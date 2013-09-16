package main

import "testing"

func TestP11(t *testing.T) {
	in, out := "0011_input.txt", 70600674
	if x := p11(in); x != out {
		t.Errorf("p11(%s) = %d, should be %d", in, x, out)
	}
}

func BenchmarkP11(b *testing.B) {
	in := "0011_input.txt"
	for i := 0; i < b.N; i++ {
		p11(in)
	}
}
