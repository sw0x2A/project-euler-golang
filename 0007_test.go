package main

import "testing"

func TestP7(t *testing.T) {
	pairs := [...][2]int{{6, 13}, {10001, 104743}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p7(in); x != out {
			t.Errorf("p7(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP7(b *testing.B) {
	in := 10001
	for i := 0; i < b.N; i++ {
		p7(in)
	}
}
