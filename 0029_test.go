package main

import "testing"

func TestP29(t *testing.T) {
	pairs := [...][2]int{{5, 15}, {100, 9183}}
	for _, pair := range pairs {
		in, out := pair[0], int(pair[1])
		if x := p29(in); x != out {
			t.Errorf("p29(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP29(b *testing.B) {
	in := 100
	for i := 0; i < b.N; i++ {
		p29(in)
	}
}
