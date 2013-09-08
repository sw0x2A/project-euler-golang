package main

import "testing"

func TestP2(t *testing.T) {
	pairs := [...][2]int{{100, 44}, {4000000, 4613732}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p2(in); x != out {
			t.Errorf("p2(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP2(b *testing.B) {
	in := 4000000
	for i := 0; i < b.N; i++ {
		p2(in)
	}
}
