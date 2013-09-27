package main

import "testing"

func TestP15(t *testing.T) {
	pairs := [...][2]uint64{{2, 6}, {20, 137846528820}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p15(in); x != out {
			t.Errorf("p15(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP15(b *testing.B) {
	in := uint64(20)
	for i := 0; i < b.N; i++ {
		p15(in)
	}
}
