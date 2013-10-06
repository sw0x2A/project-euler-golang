// https://projecteuler.net/problem=5

package main

func p5(k uint) uint {
	p := []uint{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31}
	r := uint(1)
	var d uint
	for i := range p {
		d = p[i]
		for d <= k {
			d *= p[i]
		}
		d /= p[i]
		r *= d
	}
	return r
}
