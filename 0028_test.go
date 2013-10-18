package main

import "testing"

func TestP28(t *testing.T) {
	pairs := [...][2]int{{5, 101}, {1001, 669171001}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p28(in); x != out {
			t.Errorf("p28(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP28(b *testing.B) {
	in := 1001
	for i := 0; i < b.N; i++ {
		p28(in)
	}
}
