package main

import "testing"

func TestP5(t *testing.T) {
	pairs := [...][2]uint{{10, 2520}, {20, 232792560}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p5(in); x != out {
			t.Errorf("p5(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP5(b *testing.B) {
	in := uint(20)
	for i := 0; i < b.N; i++ {
		p5(in)
	}
}
