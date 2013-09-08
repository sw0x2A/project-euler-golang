package main

import "testing"

func TestP1easy(t *testing.T) {
	pairs := [...][2]int{{9, 23}, {999, 233168}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p1easy(in); x != out {
			t.Errorf("p1easy(%d) = %d, should be %d", in, x, out)
		}
	}
}

func TestP1fast(t *testing.T) {
	pairs := [...][2]int{{9, 23}, {999, 233168}}
	for _, pair := range pairs {
		in, out := pair[0], pair[1]
		if x := p1fast(in); x != out {
			t.Errorf("p1fast(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP1easy(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		p1easy(in)
	}
}

func BenchmarkP1fast(b *testing.B) {
	in := 1000000000
	for i := 0; i < b.N; i++ {
		p1fast(in)
	}
}
