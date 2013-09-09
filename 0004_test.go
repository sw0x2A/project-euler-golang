package main

import "testing"

func TestReverse(t *testing.T) {
	pairs := [...][2]int{{111, 111}, {7337, 7337}, {906609, 906609}}
	for _, pair := range pairs {
		in, out := pair[0], int(pair[1])
		if x := reverse(in); x != out {
			t.Errorf("reverse(%d) = %d, should be %d", in, x, out)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	in := 906609
	for i := 0; i < b.N; i++ {
		reverse(in)
	}
}

func TestIsPalindrome(t *testing.T) {
	pairs := map[int]bool{1234: false, 906609: true}
	for k, _ := range pairs {
		in, out := k, pairs[k]
		if x := isPalindrome(in); x != out {
			t.Errorf("isPalindrome(%d) = %t, should be %t", in, x, out)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	in := 906609
	for i := 0; i < b.N; i++ {
		isPalindrome(in)
	}
}

func TestP4(t *testing.T) {
	out := 906609
	if x := p4(); x != out {
		t.Errorf("p4() = %d, should be %d", x, out)
	}
}

func BenchmarkP4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p4()
	}
}
