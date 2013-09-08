package main

import "testing"

func TestP3(t *testing.T) {
	pairs := [...][2]int64{{13195, 29}, {600851475143, 6857}}
	for _, pair := range pairs {
		in, out := pair[0], int(pair[1])
		if x := p3(in); x != out {
			t.Errorf("p3(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkP3(b *testing.B) {
	in := int64(600851475143)
	for i := 0; i < b.N; i++ {
		p3(in)
	}
}
