package main

import "testing"

func TestP16(t *testing.T) {
	pairs := [...][2]int{{15, 26}, {1000, 1366}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p16(in); x != out {
			t.Errorf("p16(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP16(b *testing.B) {
	in := 1000
	for i := 0; i < b.N; i++ {
		p16(in)
	}
}
