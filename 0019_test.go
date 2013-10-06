package main

import "testing"

func TestP19(t *testing.T) {
	out := 171
	if x := p19(); x != out {
		t.Errorf("p19() = %d, should be %d", x, out)
	}
}

func BenchmarkP19(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p19()
	}
}

func TestP19weekday(t *testing.T) {
	out := 171
	if x := p19weekday(); x != out {
		t.Errorf("p19weekday() = %d, should be %d", x, out)
	}
}

func BenchmarkP19weekday(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p19weekday()
	}
}
