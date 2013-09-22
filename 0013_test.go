package main

import "testing"

func TestP13(t *testing.T) {
	out := "5537376230"
	if x := p13(); x != out {
		t.Errorf("p13() = %s, should be %s", x, out)
	}
}

func BenchmarkP13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p13()
	}
}

func TestP13big(t *testing.T) {
	out := "5537376230"
	if x := p13big(); x != out {
		t.Errorf("p13big() = %s, should be %s", x, out)
	}
}

func BenchmarkP13big(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p13big()
	}
}
